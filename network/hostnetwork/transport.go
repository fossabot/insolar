/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted (subject to the limitations in the disclaimer below) provided that the following conditions are met:
 *
 *  Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *  Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *  Neither the name of Insolar Technologies nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package hostnetwork

import (
	"context"
	"fmt"
	"time"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/sequence"
	"github.com/insolar/insolar/network/transport"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/insolar/insolar/network/transport/packet"
	"github.com/insolar/insolar/network/transport/packet/types"
	"github.com/insolar/insolar/network/transport/relay"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

type hostTransport struct {
	transportBase
	handlers map[types.PacketType]network.RequestHandler
}

type packetWrapper packet.Packet

func (p *packetWrapper) GetSender() core.RecordRef {
	return p.Sender.NodeID
}

func (p *packetWrapper) GetSenderHost() *host.Host {
	return p.Sender
}

func (p *packetWrapper) GetType() types.PacketType {
	return p.Type
}

func (p *packetWrapper) GetData() interface{} {
	return p.Data
}

func (p *packetWrapper) GetRequestID() network.RequestID {
	return p.RequestID
}

type future struct {
	transport.Future
}

// Response get channel that receives response to sent request
func (f future) Response() <-chan network.Response {
	in := transport.Future(f).Result()
	out := make(chan network.Response, cap(in))
	go func(in <-chan *packet.Packet, out chan<- network.Response) {
		for packet := range in {
			out <- (*packetWrapper)(packet)
		}
		close(out)
	}(in, out)
	return out
}

// GetResponse get response to sent request with `duration` timeout
func (f future) GetResponse(duration time.Duration) (network.Response, error) {
	result, err := f.GetResult(duration)
	if err != nil {
		return nil, err
	}
	return (*packetWrapper)(result), nil
}

// GetRequest get initiating request.
func (f future) GetRequest() network.Request {
	request := transport.Future(f).Request()
	return (*packetWrapper)(request)
}

func (h *hostTransport) processMessage(msg *packet.Packet) {
	ctx, logger := inslogger.WithTraceField(context.Background(), msg.TraceID)
	logger.Debugf("Got %s request from host %s; RequestID: %d", msg.Type.String(), msg.Sender.String(), msg.RequestID)
	handler, exist := h.handlers[msg.Type]
	if !exist {
		logger.Errorf("No handler set for packet type %s from node %s",
			msg.Type.String(), msg.Sender.NodeID.String())
		return
	}
	ctx, span := instracer.StartSpan(ctx, "hostTransport.processMessage")
	span.AddAttributes(
		trace.StringAttribute("msg receiver", msg.Receiver.Address.String()),
		trace.StringAttribute("msg trace", msg.TraceID),
		trace.StringAttribute("msg type", msg.Type.String()),
	)
	defer span.End()
	response, err := handler(ctx, (*packetWrapper)(msg))
	if err != nil {
		logger.Errorf("Error handling request %s from node %s: %s",
			msg.Type.String(), msg.Sender.NodeID.String(), err)
		return
	}
	r := response.(*packetWrapper)
	err = h.transport.SendResponse(ctx, msg.RequestID, (*packet.Packet)(r))
	if err != nil {
		logger.Error(err)
	}
}

// SendRequestPacket send request packet to a remote node.
func (h *hostTransport) SendRequestPacket(ctx context.Context, request network.Request, receiver *host.Host) (network.Future, error) {
	inslogger.FromContext(ctx).Debugf("Send %s request to host %s", request.GetType().String(), receiver.String())
	f, err := h.transport.SendRequest(ctx, h.buildRequest(ctx, request, receiver))
	if err != nil {
		return nil, err
	}
	return future{Future: f}, nil
}

// RegisterPacketHandler register a handler function to process incoming request packets of a specific type.
func (h *hostTransport) RegisterPacketHandler(t types.PacketType, handler network.RequestHandler) {
	_, exists := h.handlers[t]
	if exists {
		panic(fmt.Sprintf("multiple handlers for packet type %s are not supported!", t.String()))
	}
	h.handlers[t] = handler
}

// BuildResponse create response to an incoming request with Data set to responseData.
func (h *hostTransport) BuildResponse(ctx context.Context, request network.Request, responseData interface{}) network.Response {
	sender := request.(*packetWrapper).Sender
	p := packet.NewBuilder(h.origin).Type(request.GetType()).Receiver(sender).RequestID(request.GetRequestID()).
		Response(responseData).TraceID(inslogger.TraceID(ctx)).Build()
	return (*packetWrapper)(p)
}

func NewInternalTransport(conf configuration.Configuration, nodeRef string) (network.InternalTransport, error) {
	tp, err := transport.NewTransport(conf.Host.Transport, relay.NewProxy())
	if err != nil {
		return nil, errors.Wrap(err, "error creating transport")
	}
	origin, err := getOrigin(tp, nodeRef)
	if err != nil {
		return nil, errors.Wrap(err, "error getting origin")
	}
	result := &hostTransport{handlers: make(map[types.PacketType]network.RequestHandler)}
	result.sequenceGenerator = sequence.NewGeneratorImpl()
	result.transport = tp
	result.origin = origin
	result.messageProcessor = result.processMessage
	return result, nil
}
