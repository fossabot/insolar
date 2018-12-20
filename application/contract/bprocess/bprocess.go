package bprocess

import (
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type BProcess struct {
	foundation.BaseContract
	Name string
}

func (bp *BProcess) ToJSON() ([]byte, error) {

	memberJSON, err := json.Marshal(bp)
	if err != nil {
		return nil, fmt.Errorf("[ ToJSON ]: %s", err.Error())
	}

	return memberJSON, nil
}

func New(name string) (*BProcess, error) {
	return &BProcess{
		Name: name,
	}, nil
}
