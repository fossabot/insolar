// +build functest

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

package functest

import (
	"encoding/base64"
	"encoding/json"
	"github.com/insolar/insolar/application/contract/member"
	"github.com/stretchr/testify/require"
	"testing"
)

func createOrganization(t *testing.T, name, requisites string) *user {
	member, err := newUserWithKeys()
	require.NoError(t, err)
	result, err := signedRequest(&root, "CreateOrganization", name, member.pubKey, requisites)
	require.NoError(t, err)
	ref, ok := result.(string)
	require.True(t, ok)
	member.ref = ref
	return member
}

func addMemberToOrganization(t *testing.T, memberRef, organizationRef string) string {
	result, err := signedRequest(&root, "AddMemberToOrganization", memberRef, organizationRef)
	require.NoError(t, err)
	ref, ok := result.(string)
	require.True(t, ok)
	require.NotEqual(t, "", ref)

	return ref
}

func DumpAllOrganizationMembers(t *testing.T, organizationRef string) (result []member.Member) {
	resp, err := signedRequest(&root, "DumpAllOrganizationMembers", organizationRef)
	require.NoError(t, err)
	require.NotNil(t, resp)

	data, err := base64.StdEncoding.DecodeString(resp.(string))
	require.NoError(t, err)

	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	return
}

func TestGetParticipants(t *testing.T) {
	const (
		MEMBER_NAME1 = "Member1"
		MEMBER_NAME2 = "Member2"
	)

	// Create members
	member1 := createMember(t, MEMBER_NAME1)
	member2 := createMember(t, MEMBER_NAME2)

	// Create organization
	organizationRef := createOrganization(t, "Organization", "Inn 2323232")

	// Add members
	addMemberToOrganization(t, member1.ref, organizationRef.ref)
	addMemberToOrganization(t, member2.ref, organizationRef.ref)

	// Get all members from organization
	members := DumpAllOrganizationMembers(t, organizationRef.ref)
	require.Equal(t, members[1].PublicKey, member1.pubKey)
	require.Equal(t, members[1].Name, MEMBER_NAME1)
	require.Equal(t, members[0].PublicKey, member2.pubKey)
	require.Equal(t, members[0].Name, MEMBER_NAME2)
}
