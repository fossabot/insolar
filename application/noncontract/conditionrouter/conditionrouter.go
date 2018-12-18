package conditionrouter

import (
	"github.com/insolar/insolar/application/noncontract/condition"
)

type ConditionRouter struct {
	Condition condition.Condition
}

func (cr ConditionRouter) Route() {
	//if cr.Condition.GetResult(cr.GetReference()) ...
}
