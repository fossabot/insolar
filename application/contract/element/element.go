package element

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Element struct {
	foundation.BaseContract
	Name string
}

// New creates new Element
func New(name string) (*Element, error) {
	return &Element{
		Name: name,
	}, nil
}
