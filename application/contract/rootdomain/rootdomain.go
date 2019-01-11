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

package rootdomain

import (
	"encoding/json"
	"fmt"
	bProcessContract "github.com/insolar/insolar/application/contract/bprocess"
	organizationContract "github.com/insolar/insolar/application/contract/organization"
	bProcessProxy "github.com/insolar/insolar/application/proxy/bprocess"
	"github.com/insolar/insolar/application/proxy/member"
	organizationProxy "github.com/insolar/insolar/application/proxy/organization"
	participantProxy "github.com/insolar/insolar/application/proxy/participant"
	"github.com/insolar/insolar/application/proxy/wallet"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

// RootDomain is smart contract representing entrance point to system
type RootDomain struct {
	foundation.BaseContract
	RootMember    core.RecordRef
	NodeDomainRef core.RecordRef
}

// CreateMember processes create member request
func (rd *RootDomain) CreateMember(name string, key string) (string, error) {
	if *rd.GetContext().Caller != rd.RootMember {
		return "", fmt.Errorf("[ CreateMember ] Only Root member can create members")
	}
	memberHolder := member.New(name, key)
	m, err := memberHolder.AsChild(rd.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateMember ] Can't save as child: %s", err.Error())
	}

	wHolder := wallet.New(1000)
	_, err = wHolder.AsDelegate(m.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateMember ] Can't save wallet as delegate: %s", err.Error())
	}

	participantHolder := participantProxy.New()
	_, err = participantHolder.AsDelegate(m.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateMember ] Can't save participant as delegate: %s", err.Error())
	}

	return m.GetReference().String(), nil
}

// GetRootMemberRef returns root member's reference
func (rd *RootDomain) GetRootMemberRef() (*core.RecordRef, error) {
	return &rd.RootMember, nil
}

func (rd *RootDomain) getUserInfoMap(m *member.Member) (map[string]interface{}, error) {
	w, err := wallet.GetImplementationFrom(m.GetReference())
	if err != nil {
		return nil, fmt.Errorf("[ getUserInfoMap ] Can't get implementation: %s", err.Error())
	}

	name, err := m.GetName()
	if err != nil {
		return nil, fmt.Errorf("[ getUserInfoMap ] Can't get name: %s", err.Error())
	}

	balance, err := w.GetBalance()
	if err != nil {
		return nil, fmt.Errorf("[ getUserInfoMap ] Can't get total balance: %s", err.Error())
	}
	return map[string]interface{}{
		"member": name,
		"wallet": balance,
	}, nil
}

// DumpUserInfo processes dump user info request
func (rd *RootDomain) DumpUserInfo(reference string) ([]byte, error) {
	caller := *rd.GetContext().Caller
	ref, err := core.NewRefFromBase58(reference)
	if err != nil {
		return nil, fmt.Errorf("[ DumpUserInfo ] Failed to parse reference: %s", err.Error())
	}
	if *ref != caller && caller != rd.RootMember {
		return nil, fmt.Errorf("[ DumpUserInfo ] You can dump only yourself")
	}
	m := member.GetObject(*ref)

	res, err := rd.getUserInfoMap(m)
	if err != nil {
		return nil, fmt.Errorf("[ DumpUserInfo ] Problem with making request: %s", err.Error())
	}

	return json.Marshal(res)
}

// DumpAllUsers processes dump all users request
func (rd *RootDomain) DumpAllUsers() ([]byte, error) {
	if *rd.GetContext().Caller != rd.RootMember {
		return nil, fmt.Errorf("[ DumpAllUsers ] Only root can call this method")
	}
	res := []map[string]interface{}{}
	iterator, err := rd.NewChildrenTypedIterator(member.GetPrototype())
	if err != nil {
		return nil, fmt.Errorf("[ DumpAllUsers ] Can't get children: %s", err.Error())
	}

	for iterator.HasNext() {
		cref, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllUsers ] Can't get next child: %s", err.Error())
		}

		if cref == rd.RootMember {
			continue
		}
		m := member.GetObject(cref)
		userInfo, err := rd.getUserInfoMap(m)
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllUsers ] Problem with making request: %s", err.Error())
		}
		res = append(res, userInfo)
	}
	resJSON, _ := json.Marshal(res)
	return resJSON, nil
}

var INSATTR_Info_API = true

// Info returns information about basic objects
func (rd *RootDomain) Info() (interface{}, error) {
	res := map[string]interface{}{
		"root_member": rd.RootMember.String(),
		"node_domain": rd.NodeDomainRef.String(),
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("[ Info ] Can't marshal res: %s", err.Error())
	}
	return resJSON, nil
}

// GetNodeDomainRef returns reference of NodeDomain instance
func (rd *RootDomain) GetNodeDomainRef() (core.RecordRef, error) {
	return rd.NodeDomainRef, nil
}

// NewRootDomain creates new RootDomain
func NewRootDomain() (*RootDomain, error) {
	return &RootDomain{}, nil
}

// CreateOrganization processes create organization request
func (rd *RootDomain) CreateOrganization(name string, key string, requisites string) (string, error) {
	if *rd.GetContext().Caller != rd.RootMember {
		return "", fmt.Errorf("[ CreateOrganization ] Only Root member can create organization")
	}
	organizationHolder := organizationProxy.New(name, key, requisites)
	o, err := organizationHolder.AsChild(rd.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateOrganization ] Can't save as child: %s", err.Error())
	}

	participantHolder := participantProxy.New()
	_, err = participantHolder.AsDelegate(o.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateOrganization ] Can't save participant as delegate: %s", err.Error())
	}

	return o.GetReference().String(), nil
}

// DumpAllOrganizations processes dump all organizations request
func (rd *RootDomain) DumpAllOrganizations() ([]byte, error) {
	if *rd.GetContext().Caller != rd.RootMember {
		return nil, fmt.Errorf("[ DumpAllOrganizations ] Only root can call this method")
	}

	iterator, err := rd.NewChildrenTypedIterator(organizationProxy.GetPrototype())
	if err != nil {
		return nil, fmt.Errorf("[ DumpAllOrganizations ] Can't get children: %s", err.Error())
	}

	res := []organizationContract.Organization{}
	for iterator.HasNext() {
		cref, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllOrganizations ] Can't get next child: %s", err.Error())
		}

		o := organizationProxy.GetObject(cref)

		organizationsJSON, err := o.ToJSON()

		cOrganization := organizationContract.Organization{}
		err = json.Unmarshal(organizationsJSON, &cOrganization)
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllOrganizations ] Problem with unmarshal organization from response: %s", err.Error())
		}

		res = append(res, cOrganization)
	}
	resJSON, _ := json.Marshal(res)
	return resJSON, nil
}

// CreateBProcess processes create business process request
func (rd *RootDomain) CreateBProcess(name string) (string, error) {
	if *rd.GetContext().Caller != rd.RootMember {
		return "", fmt.Errorf("[ CreateOrganization ] Only Root member can create organization")
	}
	bProcessHolder := bProcessProxy.New(name)
	bp, err := bProcessHolder.AsChild(rd.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateOrganization ] Can't save as child: %s", err.Error())
	}

	return bp.GetReference().String(), nil
}

// DumpAllBProcesses processes dump all bProcesses request
func (rd *RootDomain) DumpAllBProcesses() ([]byte, error) {

	iterator, err := rd.NewChildrenTypedIterator(bProcessProxy.GetPrototype())
	if err != nil {
		return nil, fmt.Errorf("[ DumpAllBProcesses ] Can't get children: %s", err.Error())
	}

	res := []bProcessContract.BProcess{}
	for iterator.HasNext() {
		cref, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllBProcesses ] Can't get next child: %s", err.Error())
		}

		proxyBProcess := bProcessProxy.GetObject(cref)

		bProcessJSON, err := proxyBProcess.ToJSON()

		contractBProcess := bProcessContract.BProcess{}
		err = json.Unmarshal(bProcessJSON, &contractBProcess)
		if err != nil {
			return nil, fmt.Errorf("[ DumpAllBProcesses ] Problem with unmarshal organization from response: %s", err.Error())
		}

		res = append(res, contractBProcess)
	}
	resJSON, _ := json.Marshal(res)
	return resJSON, nil
}
