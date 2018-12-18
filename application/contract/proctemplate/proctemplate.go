package proctemplate

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ProcTemplate struct {
	foundation.BaseContract
	Name string
}

func New(name string) (*ProcTemplate, error) {
	return &ProcTemplate{
		Name: name,
	}, nil
}
