package organization

import (
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
	"github.com/insolar/insolar/logicrunner/goplugin/proxyctx"
)

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = core.NewRefFromBase58("11112BMHB1MSg4MvSW62TAV8pHSiByhoauZ1EMztFAV.11111111111111111111111111111111")

// Organization holds proxy type
type Organization struct {
	Reference core.RecordRef
	Prototype core.RecordRef
	Code      core.RecordRef
}

// ContractConstructorHolder holds logic with object construction
type ContractConstructorHolder struct {
	constructorName string
	argsSerialized  []byte
}

// AsChild saves object as child
func (r *ContractConstructorHolder) AsChild(objRef core.RecordRef) (*Organization, error) {
	ref, err := proxyctx.Current.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Organization{Reference: ref}, nil
}

// AsDelegate saves object as delegate
func (r *ContractConstructorHolder) AsDelegate(objRef core.RecordRef) (*Organization, error) {
	ref, err := proxyctx.Current.SaveAsDelegate(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Organization{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref core.RecordRef) (r *Organization) {
	return &Organization{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() core.RecordRef {
	return *PrototypeReference
}

// GetImplementationFrom returns proxy to delegate of given type
func GetImplementationFrom(object core.RecordRef) (*Organization, error) {
	ref, err := proxyctx.Current.GetDelegate(object, *PrototypeReference)
	if err != nil {
		return nil, err
	}
	return GetObject(ref), nil
}

// New is constructor
func New(name string, key string, requisites string) *ContractConstructorHolder {
	var args [3]interface{}
	args[0] = name
	args[1] = key
	args[2] = requisites

	var argsSerialized []byte
	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "New", argsSerialized: argsSerialized}
}

// GetReference returns reference of the object
func (r *Organization) GetReference() core.RecordRef {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *Organization) GetPrototype() (core.RecordRef, error) {
	if r.Prototype.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 core.RecordRef
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetPrototype", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = proxyctx.Current.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Prototype = ret0
	}

	return r.Prototype, nil

}

// GetCode returns reference to the code
func (r *Organization) GetCode() (core.RecordRef, error) {
	if r.Code.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 core.RecordRef
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetCode", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = proxyctx.Current.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Code = ret0
	}

	return r.Code, nil
}

// ToJSON is proxy generated method
func (r *Organization) ToJSON() ([]byte, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "ToJSON", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// ToJSONNoWait is proxy generated method
func (r *Organization) ToJSONNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "ToJSON", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetName is proxy generated method
func (r *Organization) GetName() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetNameNoWait is proxy generated method
func (r *Organization) GetNameNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetPublicKey is proxy generated method
func (r *Organization) GetPublicKey() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetPublicKey", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetPublicKeyNoWait is proxy generated method
func (r *Organization) GetPublicKeyNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetPublicKey", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetRequisites is proxy generated method
func (r *Organization) GetRequisites() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetRequisites", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetRequisitesNoWait is proxy generated method
func (r *Organization) GetRequisitesNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetRequisites", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// VerifySig is proxy generated method
func (r *Organization) VerifySig(method string, params []byte, seed []byte, sign []byte) error {
	var args [4]interface{}
	args[0] = method
	args[1] = params
	args[2] = seed
	args[3] = sign

	var argsSerialized []byte

	ret := [1]interface{}{}
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "VerifySig", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return err
	}

	if ret0 != nil {
		return ret0
	}
	return nil
}

// VerifySigNoWait is proxy generated method
func (r *Organization) VerifySigNoWait(method string, params []byte, seed []byte, sign []byte) error {
	var args [4]interface{}
	args[0] = method
	args[1] = params
	args[2] = seed
	args[3] = sign

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "VerifySig", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// AddMember is proxy generated method
func (r *Organization) AddMember(memberReferenceStr string) (string, error) {
	var args [1]interface{}
	args[0] = memberReferenceStr

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "AddMember", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// AddMemberNoWait is proxy generated method
func (r *Organization) AddMemberNoWait(memberReferenceStr string) error {
	var args [1]interface{}
	args[0] = memberReferenceStr

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "AddMember", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetMembers is proxy generated method
func (r *Organization) GetMembers() ([]byte, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetMembers", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetMembersNoWait is proxy generated method
func (r *Organization) GetMembersNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetMembers", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}
