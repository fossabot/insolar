package bprocess

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type BProcess struct {
	foundation.BaseContract
	Name string
}

func New(name string) (*BProcess, error) {
	return &BProcess{
		Name: name,
	}, nil
}
