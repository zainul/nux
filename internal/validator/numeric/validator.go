package numeric

import (
	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/helper"
	"github.com/zainul/nux/internal/validator/usecase"
)

const (
	NonZero     = "non_zero"
	Min         = "min"
	Max         = "max"
	GreaterThan = "gte"
	LowerThan   = "lte"
)

// New ...
func New() usecase.Validator {
	errMap := make(map[string]string, 0)
	return &validator{
		ErrorCode: errMap,
	}
}

// Validator ...
type validator struct {
	ErrorCode   map[string]string
	Min         int
	Max         int
	GreaterThan int
	LowerThan   int
	NonZero     bool
}

// Validate implementation of numeric validation
func (v *validator) Validate(val interface{}) (bool, *domain.NuxError) {
	l := val.(int)

	if v.NonZero && l == 0 {
		return false, helper.FindError(v.ErrorCode[NonZero])
	}

	if l < v.Min {
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max > 0 && v.Max >= v.Min && l > v.Max {
		return false, helper.FindError(v.ErrorCode[Max])
	}

	if l < v.LowerThan {
		return false, helper.FindError(v.ErrorCode[LowerThan])
	}

	if l > v.GreaterThan {
		return false, helper.FindError(v.ErrorCode[GreaterThan])
	}

	return true, nil
}

// SetValueFromTag set value base on struct tag
func (v *validator) SetValueFromTag(field *domain.NuxTag) {
	if field != nil {

		switch field.Tag {
		case Min:
			v.Min = helper.ParseInt(field.Value)
		case Max:
			v.Max = helper.ParseInt(field.Value)
		case GreaterThan:
			v.GreaterThan = helper.ParseInt(field.Value)
		case LowerThan:
			v.LowerThan = helper.ParseInt(field.Value)
		case NonZero:
			v.NonZero = helper.ParseBool(field.Value)
		}

		v.ErrorCode[field.Tag] = field.ErrorCode
	}
}
