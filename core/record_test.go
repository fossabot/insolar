/*
 *    Copyright 2019 Insolar Technologies
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

package core_test

import (
	"testing"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/ledger/storage/jet"
	"github.com/insolar/insolar/testutils"
	base58 "github.com/jbenet/go-base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//RecordID and RecordRef serialization tests

func TestNewIDFromBase58(t *testing.T) {
	id := testutils.RandomID()
	idStr := base58.Encode(id[:])
	id2, err := core.NewIDFromBase58(idStr)
	require.NoError(t, err)

	assert.Equal(t, id, *id2)
}

func TestRecordID_String(t *testing.T) {
	id := testutils.RandomID()
	idStr := base58.Encode(id[:])

	assert.Equal(t, idStr, id.String())
}

func TestNewRefFromBase58(t *testing.T) {
	recordID := testutils.RandomID()
	domainID := testutils.RandomID()
	refStr := recordID.String() + core.RecordRefIDSeparator + domainID.String()

	expectedRef := core.NewRecordRef(domainID, recordID)
	actualRef, err := core.NewRefFromBase58(refStr)
	require.NoError(t, err)

	assert.Equal(t, expectedRef, actualRef)
}

func TestRecordRef_String(t *testing.T) {
	ref := testutils.RandomRef()
	expectedRefStr := ref.Record().String() + core.RecordRefIDSeparator + ref.Domain().String()

	assert.Equal(t, expectedRefStr, ref.String())
}

func TestRecordID_DebugString_Jet(t *testing.T) {
	j := jet.NewID(0, []byte{})
	assert.Equal(t, "[JET 0 -]", j.DebugString())

	j = jet.NewID(1, []byte{})
	assert.Equal(t, "[JET 1 0]", j.DebugString())
	j = jet.NewID(2, []byte{})
	assert.Equal(t, "[JET 2 00]", j.DebugString())

	j = jet.NewID(1, []byte{128})
	assert.Equal(t, "[JET 1 1]", j.DebugString())
	j = jet.NewID(2, []byte{192})
	assert.Equal(t, "[JET 2 11]", j.DebugString())
}


func BenchmarkRecordID_DebugString_ZeroDepth(b *testing.B) {
	jet := jet.NewID(0, []byte{})
	for n := 0; n < b.N; n++ {
		jet.DebugString()
	}
}

func BenchmarkRecordID_DebugString_Depth1(b *testing.B) {
	jet := jet.NewID(1, []byte{128})
	for n := 0; n < b.N; n++ {
		jet.DebugString()
	}
}

func BenchmarkRecordID_DebugString_Depth5(b *testing.B) {
	jet := jet.NewID(5, []byte{128})
	for n := 0; n < b.N; n++ {
		jet.DebugString()
	}
}
