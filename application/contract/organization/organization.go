package organization

import (
	"encoding/json"
	"fmt"
	memberContract "github.com/insolar/insolar/application/contract/member"
	memberProxy "github.com/insolar/insolar/application/proxy/member"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Organization struct {
	foundation.BaseContract
	Name       string
	PublicKey  string
	Requisites string
}

func (organization *Organization) ToJSON() ([]byte, error) {

	memberJSON, err := json.Marshal(organization)
	if err != nil {
		return nil, fmt.Errorf("[ ToJSON ]: %s", err.Error())
	}

	return memberJSON, nil
}

func New(name string, key string, requisites string) (*Organization, error) {
	return &Organization{
		Name:       name,
		PublicKey:  key,
		Requisites: requisites,
	}, nil
}

func (organization *Organization) GetName() (string, error) {
	return organization.Name, nil
}

var INSATTR_GetPublicKey_API = true

func (organization *Organization) GetPublicKey() (string, error) {
	return organization.PublicKey, nil
}

func (organization *Organization) GetRequisites() (string, error) {
	return organization.Requisites, nil
}

func (organization *Organization) VerifySig(method string, params []byte, seed []byte, sign []byte) error {
	args, err := core.MarshalArgs(organization.GetReference(), method, params, seed)
	if err != nil {
		return fmt.Errorf("[ verifySig ] Can't MarshalArgs: %s", err.Error())
	}
	key, err := organization.GetPublicKey()
	if err != nil {
		return fmt.Errorf("[ verifySig ]: %s", err.Error())
	}

	publicKey, err := foundation.ImportPublicKey(key)
	if err != nil {
		return fmt.Errorf("[ verifySig ] Invalid public key")
	}

	verified := foundation.Verify(args, sign, publicKey)
	if !verified {
		return fmt.Errorf("[ verifySig ] Incorrect signature")
	}
	return nil
}

// AddMemberToOrganization processes add member to organization
func (organization *Organization) AddMember(memberReferenceStr string, organizationReferenceStr string) (string, error) {

	memberReference, err := core.NewRefFromBase58(memberReferenceStr)
	if err != nil {
		return "", fmt.Errorf("[ AddMemberToOrganization ] Failed to parse member reference: %s", err.Error())
	}
	organizationReference, err := core.NewRefFromBase58(organizationReferenceStr)
	if err != nil {
		return "", fmt.Errorf("[ AddMemberToOrganization ] Failed to parse organization reference: %s", err.Error())
	}

	memberObject := memberProxy.GetObject(*memberReference)

	name, err := memberObject.GetName()
	if err != nil {
		return "", fmt.Errorf("[ AddMemberToOrganization ] Can't get name : %s", err.Error())
	}
	key, err := memberObject.GetPublicKey()
	if err != nil {
		return "", fmt.Errorf("[ AddMemberToOrganization ] Can't get key : %s", err.Error())
	}

	memberHolder := memberProxy.New(name, key)
	m, err := memberHolder.AsChild(*organizationReference)
	if err != nil {
		return "", fmt.Errorf("[ AddMemberToOrganization ] Can't save as child: %s", err.Error())
	}

	return m.GetReference().String(), nil
}

// DumpAllOrganizationMembers processes dump all organization members
func (organization *Organization) GetMembers() (resultJSON []byte, err error) {

	iterator, err := organization.NewChildrenTypedIterator(memberProxy.GetPrototype())
	if err != nil {
		return nil, fmt.Errorf("[ GetMembers ] Can't get children: %s", err.Error())
	}

	res := []memberContract.Member{}
	for iterator.HasNext() {
		cref, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("[ GetMembers ] Can't get next child: %s", err.Error())
		}

		memberProxyObject := memberProxy.GetObject(cref)

		memberJSON, err := memberProxyObject.ToJSON()
		if err != nil {
			return nil, fmt.Errorf("[ GetMembers ] Problem with making request: %s", err.Error())
		}

		memberContractObject := memberContract.Member{}
		err = json.Unmarshal(memberJSON, &memberContractObject)
		if err != nil {
			return nil, fmt.Errorf("[ GetMembers ] Problem with unmarshal member from response: %s", err.Error())
		}

		res = append(res, memberContractObject)
	}

	resultJSON, err = json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("[ GetMembers ] Problem with marshal members: %s", err.Error())
	}

	return resultJSON, nil
}
