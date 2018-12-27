package text

import (
	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/helper"
	"github.com/zainul/nux/internal/validator/usecase"
)

const (
	NonEmpty                   = "non_empty"
	Min                        = "min"
	Max                        = "max"
	ContainNumeric             = "aplhanumeric"
	Email                      = "email"
	TimeFormat                 = "time"
	TimeFormatWithCustomLayout = "timelayout"
)

// New ...
func New() usecase.Validator {
	errMap := make(map[string]string, 0)
	return &validator{
		ErrorCode: errMap,
	}
}

// validator validates for string type
type validator struct {
	ErrorCode                  map[string]string
	Min                        int  // min
	Max                        int  // max
	ContainNumeric             bool // aplhanumeric
	Email                      bool // email
	TimeFormat                 bool // time
	NonEmpty                   bool
	TimeFormatWithCustomLayout string // timelayout
}

// Validate implementation of text validation
func (v *validator) Validate(val interface{}) (bool, *domain.NuxError) {
	l := len(val.(string))

	if v.NonEmpty && l == 0 {
		return false, helper.FindError(v.ErrorCode[NonEmpty])
	}

	if l < v.Min {
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max > 0 && v.Max >= v.Min && l > v.Max {
		return false, helper.FindError(v.ErrorCode[Max])
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
		case NonEmpty:
			v.NonEmpty = helper.ParseBool(field.Value)
		}

		v.ErrorCode[field.Tag] = field.ErrorCode
	}
}
