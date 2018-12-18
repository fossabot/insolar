package response

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Response struct {
	foundation.BaseContract
	Name      string
	Agreement string
	Reason    string
	Signature []byte
}

func New(name, agreement, reason string, signature []byte) (*Response, error) {
	return &Response{
		foundation.BaseContract{},
		name,
		agreement,
		reason,
		signature,
	}, nil
}

func (r *Response) GetName() (string, error) {
	return r.Name, nil
}

func (r *Response) GetAgreement() (string, error) {
	return r.Agreement, nil
}

func (r *Response) GetReason() (string, error) {
	return r.Reason, nil
}

func (r *Response) GetSignature() ([]byte, error) {
	return r.Signature, nil
}
