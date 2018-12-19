package elemtemplate

import "github.com/insolar/insolar/logicrunner/goplugin/foundation"

type ElemTemplate struct {
	foundation.BaseContract
	Name             string
	PreviousElements []ElemTemplate
	NextElements     []ElemTemplate
}

func New(name string, previousElements, nextElements []ElemTemplate) (*ElemTemplate, error) {
	return &ElemTemplate{
		Name:             name,
		PreviousElements: previousElements,
		NextElements:     nextElements,
	}, nil
}
