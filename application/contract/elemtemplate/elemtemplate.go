package elemtemplate

import (
	"fmt"
	elemTemplateProxy "github.com/insolar/insolar/application/proxy/elemtemplate"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ElemTemplate struct {
	foundation.BaseContract
	Name                       string
	PreviousElements           []elemTemplateProxy.ElemTemplate
	NextElementTemplateSuccess []elemTemplateProxy.ElemTemplate
	NextElementTemplateFail    []elemTemplateProxy.ElemTemplate
}

func New(name string, previousElements []elemTemplateProxy.ElemTemplate, nextElementTemplateSuccess []elemTemplateProxy.ElemTemplate, nextElementTemplateFail []elemTemplateProxy.ElemTemplate) (*ElemTemplate, error) {
	return &ElemTemplate{
		Name:                       name,
		PreviousElements:           previousElements,
		NextElementTemplateSuccess: nextElementTemplateSuccess,
		NextElementTemplateFail:    nextElementTemplateFail,
	}, nil
}

func NewFromRefs(name string, previousElemTemplatesRefs []string, nextElementTemplateSuccessRefs []string, nextElementTemplateFailRefs []string) (*ElemTemplate, error) {

	previousElemTemplates, err := GetElemTemplatesByRefStrs(previousElemTemplatesRefs)
	if err != nil {
		return nil, err
	}

	nextElementTemplateSuccess, err := GetElemTemplatesByRefStrs(nextElementTemplateSuccessRefs)
	if err != nil {
		return nil, err
	}

	nextElementTemplateFail, err := GetElemTemplatesByRefStrs(nextElementTemplateFailRefs)
	if err != nil {
		return nil, err
	}

	return New(name, previousElemTemplates[:], nextElementTemplateSuccess[:], nextElementTemplateFail[:])
}

func GetElemTemplatesByRefStrs(refStrs []string) ([]elemTemplateProxy.ElemTemplate, error) {

	var elemTemplates [len(refStrs)]elemTemplateProxy.ElemTemplate

	for i, refStr := range refStrs {
		elementTemplateRef, err := core.NewRefFromBase58(refStr)
		if err != nil {
			return nil, fmt.Errorf("[ GetElemTemplatesByRefStrs ] Failed to parse element template reference: %s", err.Error())
		}

		elemTemplates[i] = *elemTemplateProxy.GetObject(*elementTemplateRef)
	}

	return elemTemplates[:], nil
}
