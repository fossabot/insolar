package bprocess

import (
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
	"github.com/insolar/insolar/logicrunner/goplugin/proxyctx"
)

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = core.NewRefFromBase58("1111HYNksMsstGtx7vLAPHoBJezLNi1Bc7ghNcTMRg.11111111111111111111111111111111")

// BProcess holds proxy type
type BProcess struct {
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
func (r *ContractConstructorHolder) AsChild(objRef core.RecordRef) (*BProcess, error) {
	ref, err := proxyctx.Current.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &BProcess{Reference: ref}, nil
}

// AsDelegate saves object as delegate
func (r *ContractConstructorHolder) AsDelegate(objRef core.RecordRef) (*BProcess, error) {
	ref, err := proxyctx.Current.SaveAsDelegate(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &BProcess{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref core.RecordRef) (r *BProcess) {
	return &BProcess{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() core.RecordRef {
	return *PrototypeReference
}

// GetImplementationFrom returns proxy to delegate of given type
func GetImplementationFrom(object core.RecordRef) (*BProcess, error) {
	ref, err := proxyctx.Current.GetDelegate(object, *PrototypeReference)
	if err != nil {
		return nil, err
	}
	return GetObject(ref), nil
}

// New is constructor
func New(name string) *ContractConstructorHolder {
	var args [1]interface{}
	args[0] = name

	var argsSerialized []byte
	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "New", argsSerialized: argsSerialized}
}

// GetReference returns reference of the object
func (r *BProcess) GetReference() core.RecordRef {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *BProcess) GetPrototype() (core.RecordRef, error) {
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
func (r *BProcess) GetCode() (core.RecordRef, error) {
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
func (r *BProcess) ToJSON() ([]byte, error) {
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
func (r *BProcess) ToJSONNoWait() error {
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

// GetProcTemplates is proxy generated method
func (r *BProcess) GetProcTemplates() ([]byte, error) {
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

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetProcTemplates", argsSerialized, *PrototypeReference)
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

// GetProcTemplatesNoWait is proxy generated method
func (r *BProcess) GetProcTemplatesNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetProcTemplates", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// CreateDocType is proxy generated method
func (r *BProcess) CreateDocType(name string, fields []docTypeProxy.Field, attachments []docTypeProxy.Attachment) (string, error) {
	var args [3]interface{}
	args[0] = name
	args[1] = fields
	args[2] = attachments

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

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "CreateDocType", argsSerialized, *PrototypeReference)
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

// CreateDocTypeNoWait is proxy generated method
func (r *BProcess) CreateDocTypeNoWait(name string, fields []docTypeProxy.Field, attachments []docTypeProxy.Attachment) error {
	var args [3]interface{}
	args[0] = name
	args[1] = fields
	args[2] = attachments

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "CreateDocType", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetDocTypes is proxy generated method
func (r *BProcess) GetDocTypes() ([]byte, error) {
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

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "GetDocTypes", argsSerialized, *PrototypeReference)
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

// GetDocTypesNoWait is proxy generated method
func (r *BProcess) GetDocTypesNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "GetDocTypes", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// СreateProcTemplate is proxy generated method
func (r *BProcess) СreateProcTemplate(name string) (string, error) {
	var args [1]interface{}
	args[0] = name

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

	res, err := proxyctx.Current.RouteCall(r.Reference, true, "СreateProcTemplate", argsSerialized, *PrototypeReference)
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

// СreateProcTemplateNoWait is proxy generated method
func (r *BProcess) СreateProcTemplateNoWait(name string) error {
	var args [1]interface{}
	args[0] = name

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, "СreateProcTemplate", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}
