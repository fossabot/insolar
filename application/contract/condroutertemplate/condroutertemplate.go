package condroutertemplate

import (
	"github.com/insolar/insolar/application/contract/condroutertemplate/condition"
	"github.com/insolar/insolar/application/proxy/element"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ConditionRouterTemplate struct {
	foundation.BaseContract
	Condition                  condition.Condition
	nextElementTemplateSuccess []element.Element
	nextElementTemplateFail    []element.Element
}

func New(condition condition.Condition) (*ConditionRouterTemplate, error) {
	return &ConditionRouterTemplate{
		Condition: condition,
	}, nil
}
