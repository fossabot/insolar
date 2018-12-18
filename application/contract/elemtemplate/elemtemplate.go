package elemtemplate

import (
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ElemTemplate struct {
	foundation.BaseContract
	Name        string
	ElementBody []byte
}
