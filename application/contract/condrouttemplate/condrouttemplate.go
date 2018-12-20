package condrouttemplate

import (
	"github.com/insolar/insolar/application/contract/condrouttemplate/condition"
	"github.com/insolar/insolar/application/contract/elemtemplate"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type CondRoutTemplate struct {
	foundation.BaseContract
	elemtemplate.ElemTemplate
	Condition condition.Condition
}

//
//func (cr CondRoutTemplate) Route() {
//	//if cr.Condition.GetResult(cr.GetReference()) ...
//}
