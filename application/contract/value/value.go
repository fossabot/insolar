package value

import (
	"encoding/binary"
	"fmt"
	"github.com/insolar/insolar/core"
	"time"
)

type ValueType string
type ArithmeticOperation string

const (
	IntegerType ValueType = "Integer"
	DateType    ValueType = "Date"

	SumOperation         ArithmeticOperation = "+"
	CompositionOperation ArithmeticOperation = "*"
	DifferenceOperation  ArithmeticOperation = "-"
	QuotientOperation    ArithmeticOperation = "/"
)

type Value interface {
	GetResult(ref core.RecordRef) []byte
}

type ArithmeticExpression struct {
	Type       ValueType
	Operation  ArithmeticOperation
	LeftValue  Value
	RightValue Value
}

func ToInt(ref core.RecordRef, v Value) (u uint64, err error) {
	return binary.BigEndian.Uint64(v.GetResult(ref)), nil
}
func ToDate(ref core.RecordRef, v Value) (t time.Time, err error) {
	return time.Parse(time.UTC.String(), string(v.GetResult(ref)))
}

func (e ArithmeticExpression) GetResult(ref core.RecordRef) (result []byte, err error) {
	switch e.Type {
	case IntegerType:

		l, err := ToInt(ref, e.LeftValue)
		if err != nil {
			return nil, err
		}
		r, err := ToInt(ref, e.RightValue)
		if err != nil {
			return nil, err
		}

		switch e.Operation {
		case SumOperation:
			binary.BigEndian.PutUint64(result, l+r)
		case CompositionOperation:
			binary.BigEndian.PutUint64(result, l*r)
		case DifferenceOperation:
			binary.BigEndian.PutUint64(result, l-r)
		case QuotientOperation:
			binary.BigEndian.PutUint64(result, l/r)
		}

		if result == nil {
			return nil, fmt.Errorf("[ GetResult ] Not valid ArithmeticExpression operation %s", e.Operation)
		} else {
			return result, nil
		}
	default:
		return nil, fmt.Errorf("[ GetResult ] Not valid ArithmeticExpression type %s", e.Type)
	}
}

type Field struct {
	DocIndex   int
	FieldIndex int
}

func (e Field) GetResult(ref core.RecordRef) (result []byte, err error) {
	//todo  get docs by ref
	return nil, nil
}

type Constant []byte
