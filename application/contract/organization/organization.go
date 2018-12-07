package organization

import (
	"fmt"
	"github.com/insolar/insolar/application/noncontract/group"
	"github.com/insolar/insolar/application/noncontract/participant"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Organization struct {
	foundation.BaseContract
	group.Group
	Requisites string
}

func New(name string, key string, requisites string) (*Organization, error) {
	return &Organization{
		foundation.BaseContract{},
		group.Group{
			participant.Participant{name, key},
		},
		requisites}, nil
}

///////////////////impl/////////////////////
func (o *Organization) GetName() (string, error) {
	return o.Participant.GetName()
}

var INSATTR_GetPublicKey_API = true

func (o *Organization) GetPublicKey() (string, error) {
	return o.Participant.GetPublicKey()
}

///////////////////impl end//////////////////

func (o *Organization) GetRequisites() (string, error) {
	return o.Requisites, nil
}

func (o *Organization) VerifySig(method string, params []byte, seed []byte, sign []byte) error {
	args, err := core.MarshalArgs(o.GetReference(), method, params, seed)
	if err != nil {
		return fmt.Errorf("[ verifySig ] Can't MarshalArgs: %s", err.Error())
	}
	key, err := o.GetPublicKey()
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
