package condrouttemplate

import (
	"github.com/insolar/insolar/application/contract/condrouttemplate/condition"
	"github.com/insolar/insolar/application/proxy/element"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type CondRoutTemplate struct {
	foundation.BaseContract
	Condition                  condition.Condition
	nextElementTemplateSuccess []element.Element
	nextElementTemplateFail    []element.Element
}

func New(condition condition.Condition) (*CondRoutTemplate, error) {
	return &CondRoutTemplate{
		Condition: condition,
	}, nil
}
