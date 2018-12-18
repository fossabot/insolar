package stage

import (
	"github.com/insolar/insolar/application/contract/elemtemplate"
	"github.com/insolar/insolar/application/contract/response"
	"github.com/insolar/insolar/application/noncontract/participant"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type DocPermission string

const (
	RWType DocPermission = "Read/Write"
	WType  DocPermission = "Write"
	RType  DocPermission = "Read"
	NType  DocPermission = "None"
)

type Stage struct {
	foundation.BaseContract
	elemtemplate.ElemTemplate
	Participant     participant.Participant
	DocsPermissions [][]DocPermission
	Response        response.Response
	ExpirationDate  string
}

func New(name string, participant participant.Participant, docsPermissions [][]DocPermission, response response.Response, expirationDate string) (*Stage, error) {
	return &Stage{
		foundation.BaseContract{},
		elemtemplate.ElemTemplate{
			Name: name,
		},
		participant,
		docsPermissions,
		response,
		expirationDate,
	}, nil
}
