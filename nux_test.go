package nux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructString struct {
	Min      string `validate:"string,min=2=E01"`
	Max      string `validate:"string,max=5=E03"`
	NonEmpty string `validate:"string,non_empty=true=E04"`
	Email    string `validate:"email,validemail=true=E05"`
}

var errs map[string]string

func InitError() {
	errorMap := make(map[string]string, 0)
	errorMap["E01"] = "Min value int should be 1"
	errorMap["E02"] = "Max value  should be 100"
	errorMap["E03"] = "Max value for string should be 10"
	errorMap["E04"] = "Cannot be empty"
	errorMap["E05"] = "Must be valid email format"
	errorMap["E06"] = "Greater than error"
	errs = errorMap
}
func TestStringMinMaxEmail(t *testing.T) {
	InitError()
	NewError(errs)

	ss := StructString{
		Max:      "ulalala",
		Email:    "email",
		Min:      "u",
		NonEmpty: "",
	}
	nuxErrors := ValidateStruct(ss)

	assert.Equal(t, 4, len(nuxErrors), "total of error should be 4")
	assert.Equal(t, "E01", nuxErrors[0].Code, "Error ordering shoud be in struct field order, 1st is E01")
	assert.Equal(t, "E03", nuxErrors[1].Code, "Error ordering shoud be in struct field order, 2nd is E03")
	assert.Equal(t, "E04", nuxErrors[2].Code, "Error ordering shoud be in struct field order, 3rd is E04")
	assert.Equal(t, "E05", nuxErrors[3].Code, "Error ordering shoud be in struct field order, 4th is E05")
}
