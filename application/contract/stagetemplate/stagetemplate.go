package stagetemplate

import (
	"fmt"
	documentProxy "github.com/insolar/insolar/application/proxy/document"
	"github.com/insolar/insolar/application/proxy/participant"
	"github.com/insolar/insolar/application/proxy/response"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type AccessModificator string

const (
	RWType AccessModificator = "Read/Write"
	WType  AccessModificator = "Write"
	RType  AccessModificator = "Read"
	NType  AccessModificator = "None"
)

type StageTemplate struct {
	foundation.BaseContract
	Participant     participant.Participant
	DocsPermissions map[documentProxy.Document]AccessModificator
	Response        response.Response
	ExpirationDate  string
}

func New(participant participant.Participant, expirationDate string) (*StageTemplate, error) {
	return &StageTemplate{
		Participant:    participant,
		ExpirationDate: expirationDate,
	}, nil
}

func (stageTemplate *StageTemplate) setDocsPermissions(docsPermissionStrs map[string]string) error {
	docsPermissions := make(map[documentProxy.Document]AccessModificator)

	for documentReferenceStr, accessModificatorStr := range docsPermissionStrs {

		documentReference, err := core.NewRefFromBase58(documentReferenceStr)
		if err != nil {
			return fmt.Errorf("[ setDocsPermissions ] Failed to parse reference: %s", err.Error())
		}

		documentObject := documentProxy.GetObject(*documentReference)

		accessModificator := AccessModificator(accessModificatorStr)

		if accessModificator == "" {
			return fmt.Errorf("[ setDocsPermissions ] Failed to parse access Modificator: %s", err.Error())
		}

		docsPermissions[*documentObject] = accessModificator
	}

	stageTemplate.DocsPermissions = docsPermissions

	return nil
}
