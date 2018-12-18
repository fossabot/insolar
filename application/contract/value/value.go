package value

import (
	"encoding/binary"
	"fmt"
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
	GetResult() []byte
}

type ArithmeticExpression struct {
	Type       ValueType
	Operation  ArithmeticOperation
	LeftValue  Value
	RightValue Value
}

func ToInt(v Value) (u uint64, err error) {
	return binary.BigEndian.Uint64(v.GetResult()), nil
}
func ToDate(v Value) (t time.Time, err error) {
	return time.Parse(time.UTC.String(), string(v.GetResult()))
}

func (e ArithmeticExpression) GetResult() (result []byte, err error) {
	switch e.Type {
	case IntegerType:

		l, err := ToInt(e.LeftValue)
		if err != nil {
			return nil, err
		}
		r, err := ToInt(e.RightValue)
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

func (e Field) GetResult() (result []byte, err error) {
	return nil, nil
}

type Constant []byte
