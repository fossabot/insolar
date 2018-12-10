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

package jet

import (
	"bytes"

	"github.com/insolar/insolar/core"
	"github.com/ugorji/go/codec"
)

// Jet contain jet record.
type Jet struct {
	ID     core.RecordID
	Prefix []byte

	Left  *Jet
	Right *Jet
}

// Find returns jet for provided reference.
func (j *Jet) Find(val []byte, pulse core.PulseNumber, depth uint) *Jet {
	if j == nil || val == nil {
		return nil
	}

	if getBit(val, depth) {
		if j.Right != nil && j.Right.ID.Pulse() <= pulse {
			return j.Right.Find(val, pulse, depth+1)
		}
	} else {
		if j.Left != nil && j.Left.ID.Pulse() <= pulse {
			return j.Left.Find(val, pulse, depth+1)
		}
	}
	return j
}

// Tree stores jet in a binary tree.
type Tree struct {
	Head *Jet
}

// Find returns jet for provided reference.
func (t *Tree) Find(val []byte, pulse core.PulseNumber) *Jet {
	return t.Head.Find(val, pulse, 0)
}

// Bytes serializes pulse.
func (t *Tree) Bytes() []byte {
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, &codec.CborHandle{})
	enc.MustEncode(t)
	return buf.Bytes()
}

func getBit(value []byte, index uint) bool {
	if index > uint(len(value)*8) {
		panic("index overflow")
	}
	byteIndex := uint(index / 8)
	bitIndex := 7 - index%8
	mask := byte(1 << bitIndex)
	return value[byteIndex]&mask != 0
}