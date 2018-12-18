package condition

import (
	"fmt"
	"github.com/insolar/insolar/application/contract/value"
)

type LogicOperation string
type ComparisonOperation string

const (
	AndOperation LogicOperation = "And"
	OrOperation  LogicOperation = "Or"

	LessOperation  ComparisonOperation = "<"
	MoreOperation  ComparisonOperation = ">"
	EqualOperation ComparisonOperation = "="
)

type Condition interface {
	GetResult() bool
}

type LogicCondition struct {
	Operation  LogicOperation
	LeftValue  Condition
	RightValue Condition
}

func (c LogicCondition) GetResult() (result bool) {
	switch c.Operation {
	case AndOperation:
		result = c.LeftValue.GetResult() && c.RightValue.GetResult()
	case OrOperation:
		result = c.LeftValue.GetResult() || c.RightValue.GetResult()
	default:
		result = false
	}

	return result
}

type ComparisonCondition struct {
	Operation  ComparisonOperation
	Type       value.ValueType
	LeftValue  value.Value
	RightValue value.Value
}

func (c ComparisonCondition) GetResult() (result bool, err error) {
	switch c.Type {

	case value.IntegerType:

		l, err := value.ToInt(c.LeftValue)
		if err != nil {
			return false, err
		}
		r, err := value.ToInt(c.RightValue)
		if err != nil {
			return false, err
		}

		switch c.Operation {
		case LessOperation:
			return l < r, nil
		case MoreOperation:
			return l > r, nil
		case EqualOperation:
			return l == r, nil
		default:
			return false, fmt.Errorf("[ GetResult ] Not valid ComparisonCondition operation %s", c.Operation)
		}

	case value.DateType:

		l, err := value.ToDate(c.LeftValue)
		if err != nil {
			return false, err
		}
		r, err := value.ToDate(c.RightValue)
		if err != nil {
			return false, err
		}

		switch c.Operation {
		case LessOperation:
			return l.Before(r), nil
		case MoreOperation:
			return l.After(r), nil
		case EqualOperation:
			return l.Equal(r), nil
		default:
			return false, fmt.Errorf("[ GetResult ] Not valid ComparisonCondition operation %s", c.Operation)
		}

	default:
		return false, fmt.Errorf("[ GetResult ] Not valid ComparisonCondition type %s", c.Type)
	}
}
