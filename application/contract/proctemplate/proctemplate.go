package proctemplate

import (
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/application/contract/condroutertemplate/condition"
	documentContract "github.com/insolar/insolar/application/contract/document"
	"github.com/insolar/insolar/application/proxy/allowance"
	condRouterTemplateProxy "github.com/insolar/insolar/application/proxy/condroutertemplate"
	docTypeProxy "github.com/insolar/insolar/application/proxy/doctype"
	documentProxy "github.com/insolar/insolar/application/proxy/document"
	elemTemplateProxy "github.com/insolar/insolar/application/proxy/elemtemplate"
	participantProxy "github.com/insolar/insolar/application/proxy/participant"
	stageTemplateProxy "github.com/insolar/insolar/application/proxy/stagetemplate"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type ProcTemplate struct {
	foundation.BaseContract
	Name string
}

func (procTemplate *ProcTemplate) ToJSON() ([]byte, error) {

	memberJSON, err := json.Marshal(procTemplate)
	if err != nil {
		return nil, fmt.Errorf("[ ToJSON ]: %s", err.Error())
	}

	return memberJSON, nil
}

func New(name string) (*ProcTemplate, error) {
	return &ProcTemplate{
		Name: name,
	}, nil
}

// CreateDocument processes create document request
func (procTemplate *ProcTemplate) CreateDocument(name string, docTypeReferenceStr string) (string, error) {

	docTypeReference, err := core.NewRefFromBase58(docTypeReferenceStr)
	if err != nil {
		return "", fmt.Errorf("[ CreateDocument ] Failed to parse document type reference: %s", err.Error())
	}
	docTypeObject := *docTypeProxy.GetObject(*docTypeReference)

	documentHolder := documentProxy.New(name, docTypeObject)

	d, err := documentHolder.AsChild(procTemplate.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateDocument ] Can't save as child: %s", err.Error())
	}

	return d.GetReference().String(), nil
}

// GetDocuments processes dump all documents
func (procTemplate *ProcTemplate) GetDocuments() (resultJSON []byte, err error) {
	iterator, err := procTemplate.NewChildrenTypedIterator(allowance.GetPrototype())
	if err != nil {
		return nil, fmt.Errorf("[ GetBalance ] Can't get children: %s", err.Error())
	}

	res := []documentContract.Document{}
	for iterator.HasNext() {
		cref, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("[ GetDocTypes ] Can't get next child: %s", err.Error())
		}

		documentProxyObject := documentProxy.GetObject(cref)

		documentJSON, err := documentProxyObject.ToJSON()
		if err != nil {
			return nil, fmt.Errorf("[ GetDocTypes ] Problem with making request: %s", err.Error())
		}

		documentContractObject := documentContract.Document{}
		err = json.Unmarshal(documentJSON, &documentContractObject)
		if err != nil {
			return nil, fmt.Errorf("[ GetDocTypes ] Problem with unmarshal children from response: %s", err.Error())
		}

		res = append(res, documentContractObject)
	}

	resultJSON, err = json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("[ GetDocTypes ] Problem with marshal children: %s", err.Error())
	}

	return resultJSON, nil
}

func (procTemplate *ProcTemplate) createElementTemplate(
	name string,
	previousElemTemplatesRefs []string,
	nextElementTemplateSuccess []string,
	nextElementTemplateFail []string) {
}

// CreateStageTemplate processes create stage request
func (procTemplate *ProcTemplate) CreateStageTemplate(
	name string,
	previousElemTemplatesRefs []string,
	nextElementTemplateSuccessRefs []string,
	nextElementTemplateFailRefs []string,
	participantsRef string,
	expirationDate string) (string, error) {

	elemTemplateHolder := elemTemplateProxy.NewFromRefs(name, previousElemTemplatesRefs, nextElementTemplateSuccessRefs, nextElementTemplateFailRefs)
	et, err := elemTemplateHolder.AsChild(procTemplate.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateStageTemplate ] Can't save as child: %s", err.Error())
	}

	ref, err := core.NewRefFromBase58(participantsRef)
	if err != nil {
		return "", fmt.Errorf("[ CreateStageTemplate ] Failed to parse participant reference: %s", err.Error())
	}

	participantObject := *participantProxy.GetObject(*ref)

	stageTemplateHolder := stageTemplateProxy.New(participantObject, expirationDate)
	st, err := stageTemplateHolder.AsChild(et.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateStageTemplate ] Can't save as child: %s", err.Error())
	}

	return st.GetReference().String(), nil
}

// CreateConditionRouterTemplate processes create Condition Router template request
func (procTemplate *ProcTemplate) CreateConditionRouterTemplate(name string,
	previousElemTemplatesRefs []string,
	nextElementTemplateSuccessRefs []string,
	nextElementTemplateFailRefs []string,
	conditionJSON []byte) (string, error) {

	elemTemplateHolder := elemTemplateProxy.NewFromRefs(name, previousElemTemplatesRefs, nextElementTemplateSuccessRefs, nextElementTemplateFailRefs)
	et, err := elemTemplateHolder.AsChild(procTemplate.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateConditionRouterTemplate ] Can't save as child: %s", err.Error())
	}

	var conditionObject condition.Condition

	json.Unmarshal(conditionJSON, &conditionObject)

	condRouterTemplateHolder := condRouterTemplateProxy.New(conditionObject)
	st, err := condRouterTemplateHolder.AsChild(et.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CreateConditionRouterTemplate ] Can't save as child: %s", err.Error())
	}

	return st.GetReference().String(), nil
}
