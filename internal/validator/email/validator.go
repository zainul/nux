package email

import (
	"regexp"

	"github.com/zainul/nux/internal/helper"

	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/validator/usecase"
)

var mailRegex = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

const (
	// TypeEmail is tag for email validation
	TypeEmail = "validemail"
)

// New ...
func New() usecase.Validator {
	errMap := make(map[string]string, 0)
	return &validator{
		ErrorCode: errMap,
	}
}

// validator ...
type validator struct {
	ErrorCode map[string]string
	TypeEmail string
}

// *validate implementation of decimal validation
func (v *validator) Validate(val interface{}) (bool, *domain.NuxError) {
	if !mailRegex.MatchString(val.(string)) {
		return false, helper.FindError(v.ErrorCode[TypeEmail])
	}
	return true, nil
}

// SetValueFromTag set value base on struct tag
func (v *validator) SetValueFromTag(field *domain.NuxTag) {
	if field != nil {
		switch field.Tag {
		case TypeEmail:
			v.TypeEmail = field.Value
		}
		v.ErrorCode[field.Tag] = field.ErrorCode
	}
}
