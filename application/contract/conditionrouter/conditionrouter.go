package conditionrouter

import (
	"github.com/insolar/insolar/application/noncontract/condition"
	"github.com/insolar/insolar/application/noncontract/elementtemplate"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Stage struct {
	foundation.BaseContract
	elementtemplate.ElementTemplate
	Condition condition.Condition
}

func New(name string, condition condition.Condition) (*Stage, error) {
	return &Stage{
		foundation.BaseContract{},
		elementtemplate.ElementTemplate{
			Name: name,
		},
		condition,
	}, nil
}
