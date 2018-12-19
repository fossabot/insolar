package stagetemplate

import (
	"github.com/insolar/insolar/application/contract/elemtemplate"
	"github.com/insolar/insolar/application/contract/participant"
	"github.com/insolar/insolar/application/contract/response"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type DocPermission string

const (
	RWType DocPermission = "Read/Write"
	WType  DocPermission = "Write"
	RType  DocPermission = "Read"
	NType  DocPermission = "None"
)

type StageTemplate struct {
	foundation.BaseContract
	elemtemplate.ElemTemplate
	Participants    []participant.Participant
	DocsPermissions [][]DocPermission
	Response        response.Response
	ExpirationDate  string
}

func New(name string, participants []participant.Participant, docsPermissions [][]DocPermission, response response.Response, expirationDate string) (*StageTemplate, error) {
	return &StageTemplate{
		foundation.BaseContract{},
		elemtemplate.ElemTemplate{
			Name: name,
		},
		participants,
		docsPermissions,
		response,
		expirationDate,
	}, nil
}
