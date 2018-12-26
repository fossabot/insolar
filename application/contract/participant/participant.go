package participant

import (
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Participant struct {
	foundation.BaseContract
}

func (participant *Participant) ToJSON() ([]byte, error) {

	memberJSON, err := json.Marshal(participant)
	if err != nil {
		return nil, fmt.Errorf("[ ToJSON ]: %s", err.Error())
	}

	return memberJSON, nil
}

func New() (*Participant, error) {
	return &Participant{}, nil
}
