package conditionrouter

import (
	"github.com/insolar/insolar/application/contract/elemtemplate"
	"github.com/insolar/insolar/application/noncontract/condition"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ConditionRouter struct {
	foundation.BaseContract
	elemtemplate.ElemTemplate
	Condition condition.Condition
}

//
//func (cr ConditionRouter) Route() {
//	//if cr.Condition.GetResult(cr.GetReference()) ...
//}
