/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package core

import (
	"crypto/ecdsa"
)

// NodeState is the state of the node
type NodeState uint8

// TODO: document all node states
const (
	// Joined
	NodeJoined = NodeState(iota + 1)
	// Prepared
	NodePrepared
	// Active
	NodeActive
	// Leaved
	NodeLeaved
	// Suspended
	NodeSuspended
)

type ActiveNode struct {
	// NodeID is the unique identifier of the node
	NodeID RecordRef
	// PulseNum is the pulse number after which the new state is assigned to the node
	PulseNum PulseNumber
	// State is the node state
	State NodeState
	// JetRoles is the set of candidate JetRoles for the node
	Roles []NodeRole
	// PublicKey is the public key of the node
	PublicKey *ecdsa.PublicKey
	// Addess is the network adress of the node
	Address string
}

type ActiveNodeComponent interface {
	// GetSelf get active node for the current insolard. Returns nil if the current insolard is not an active node.
	GetSelf() *ActiveNode
	// GetActiveNode get active node by its reference. Returns nil if node is not found.
	GetActiveNode(ref RecordRef) *ActiveNode
	// GetActiveNodes get active nodes.
	GetActiveNodes() []*ActiveNode
	// GetActiveNodesByRole get active nodes by role
	GetActiveNodesByRole(role JetRole) []RecordRef
}
