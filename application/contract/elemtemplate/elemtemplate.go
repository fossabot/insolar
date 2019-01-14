package elemtemplate

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ElemTemplate struct {
	foundation.BaseContract
	Name                       string
	PreviousElements           []ElemTemplate
	NextElementTemplateSuccess []ElemTemplate
	NextElementTemplateFail    []ElemTemplate
}

func New(name string, previousElements []ElemTemplate, nextElementTemplateSuccess []ElemTemplate, nextElementTemplateFail []ElemTemplate) (*ElemTemplate, error) {
	return &ElemTemplate{
		Name:                       name,
		PreviousElements:           previousElements,
		NextElementTemplateSuccess: nextElementTemplateSuccess,
		NextElementTemplateFail:    nextElementTemplateFail,
	}, nil
}
