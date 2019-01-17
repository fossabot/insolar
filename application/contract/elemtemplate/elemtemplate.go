package elemtemplate

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ElemTemplate struct {
	foundation.BaseContract
	Name                        string
	PreviousElemTemplateRefs    []string
	NextElemTemplateSuccessRefs []string
	NextElemTemplateFailRefs    []string
}

func New(name string, previousElements []string, nextElementTemplateSuccess []string, nextElementTemplateFail []string) (*ElemTemplate, error) {
	return &ElemTemplate{
		Name:                        name,
		PreviousElemTemplateRefs:    previousElements,
		NextElemTemplateSuccessRefs: nextElementTemplateSuccess,
		NextElemTemplateFailRefs:    nextElementTemplateFail,
	}, nil
}

func (elemTemplate *ElemTemplate) SetPreviousElemTemplateRef(previousElemTemplateRef string) error {
	elemTemplate.PreviousElemTemplateRefs = append(elemTemplate.PreviousElemTemplateRefs, previousElemTemplateRef)
	return nil
}

func (elemTemplate *ElemTemplate) SetNextElemTemplateSuccessRef(nextElemTemplateSuccessRef string) error {
	elemTemplate.NextElemTemplateSuccessRefs = append(elemTemplate.NextElemTemplateSuccessRefs, nextElemTemplateSuccessRef)
	return nil
}

func (elemTemplate *ElemTemplate) SetNextElemTemplateFailRef(nextElemTemplateFailRef string) error {
	elemTemplate.NextElemTemplateFailRefs = append(elemTemplate.NextElemTemplateFailRefs, nextElemTemplateFailRef)
	return nil
}
