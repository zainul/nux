package internal

import (
	"reflect"
	"strings"

	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/helper"
	"github.com/zainul/nux/internal/validator/email"
	"github.com/zainul/nux/internal/validator/errors"
	"github.com/zainul/nux/internal/validator/numeric"
	"github.com/zainul/nux/internal/validator/text"
)

// Name of the struct tag used in examples.
const tagName = "validate"

// Validator is generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(
		interface{}, // input
	) (bool, *domain.NuxError)
	SetValueFromTag(*domain.NuxTag)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

// Validate ...
func (v DefaultValidator) Validate(val interface{}) (bool, *domain.NuxError) {
	return true, nil
}

// SetValueFromTag ...
func (v DefaultValidator) SetValueFromTag(*domain.NuxTag) {

}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case "number":
		val := numeric.New()

		for _, validation := range args[1:] {
			val.SetValueFromTag(helper.SplitError(validation))
		}
		return val
	case "string":
		val := text.New()

		for _, validation := range args[1:] {
			val.SetValueFromTag(helper.SplitError(validation))
		}
		return val
	case "email":
		val := email.New()

		for _, validation := range args[1:] {
			val.SetValueFromTag(helper.SplitError(validation))
		}
		return val
	}

	return DefaultValidator{}
}

// ValidateStruct is Performs actual data validation using validator definitions on the struct
func ValidateStruct(s interface{}) []domain.NuxError {
	errs := make([]domain.NuxError, 0)

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			errs = append(errs, *err)
		}
	}

	return errs
}

// NewError ...
func NewError(errs map[string]string) {
	errors.InitializeError(errs)
}
