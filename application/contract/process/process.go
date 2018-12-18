package process

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Process struct {
	foundation.BaseContract
	Name string
}

func New(name string) (*Process, error) {
	return &Process{
		Name: name,
	}, nil
}
