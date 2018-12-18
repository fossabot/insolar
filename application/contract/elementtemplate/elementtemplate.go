package elementtemplate

import (
	"github.com/insolar/insolar/application/noncontract/conditionrouter"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ElementTemplate struct {
	foundation.BaseContract
	Name            string
	ConditionRouter conditionrouter.ConditionRouter
}
