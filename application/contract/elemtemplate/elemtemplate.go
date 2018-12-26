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
