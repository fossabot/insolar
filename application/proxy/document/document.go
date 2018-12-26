package document

import (
	"github.com/insolar/insolar/application/proxy/doctype"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
	"github.com/insolar/insolar/logicrunner/goplugin/proxyctx"
)

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = core.NewRefFromBase58("11113W1t7CEuUK7aSThV66HFgtCFBwqhU21zgjpHHQH.11111111111111111111111111111111")

// Document holds proxy type
type Document struct {
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
func (r *ContractConstructorHolder) AsChild(objRef core.RecordRef) (*Document, error) {
	ref, err := proxyctx.Current.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Document{Reference: ref}, nil
}

// AsDelegate saves object as delegate
func (r *ContractConstructorHolder) AsDelegate(objRef core.RecordRef) (*Document, error) {
	ref, err := proxyctx.Current.SaveAsDelegate(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Document{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref core.RecordRef) (r *Document) {
	return &Document{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() core.RecordRef {
	return *PrototypeReference
}

// GetImplementationFrom returns proxy to delegate of given type
func GetImplementationFrom(object core.RecordRef) (*Document, error) {
	ref, err := proxyctx.Current.GetDelegate(object, *PrototypeReference)
	if err != nil {
		return nil, err
	}
	return GetObject(ref), nil
}

// New is constructor
func New(name string, docType doctype.DocType) *ContractConstructorHolder {
	var args [2]interface{}
	args[0] = name
	args[1] = docType

	var argsSerialized []byte
	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "New", argsSerialized: argsSerialized}
}

// GetReference returns reference of the object
func (r *Document) GetReference() core.RecordRef {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *Document) GetPrototype() (core.RecordRef, error) {
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
func (r *Document) GetCode() (core.RecordRef, error) {
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
func (r *Document) ToJSON() ([]byte, error) {
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
func (r *Document) ToJSONNoWait() error {
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
