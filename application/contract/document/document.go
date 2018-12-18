package document

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Document struct {
	foundation.BaseContract

	Name string
}

func New(name string) (*Document, error) {
	return &Document{
		Name: name,
	}, nil
}
