package decimal

import "github.com/zainul/nux/internal/domain"

// validator ...
type validator struct {
	ErrorCode   map[string]string
	Min         int
	Max         int
	GreaterThan int
	LowerThan   int
}

// Validate implementation of decimal validation
func (v *validator) Validate(val interface{}) (bool, *domain.NuxError) {
	return true, nil
}

// SetValueFromTag set value base on struct tag
func (v *validator) SetValueFromTag(*domain.NuxTag) {

}
