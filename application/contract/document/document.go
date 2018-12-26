package document

import (
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/application/proxy/doctype"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Document struct {
	foundation.BaseContract

	Name    string
	DocType doctype.DocType
}

func New(name string, docType doctype.DocType) (*Document, error) {
	return &Document{
		Name:    name,
		DocType: docType,
	}, nil
}

func (document *Document) ToJSON() ([]byte, error) {

	documentJSON, err := json.Marshal(document)
	if err != nil {
		return nil, fmt.Errorf("[ ToJSON ]: %s", err.Error())
	}

	return documentJSON, nil
}
