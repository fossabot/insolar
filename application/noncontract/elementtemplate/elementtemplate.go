package elementtemplate

import "github.com/insolar/insolar/application/contract/document"

type ElementTemplate struct {
	Name      string
	Documents []document.Document
}
