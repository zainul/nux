package nux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringStruct struct {
	Min      string `validate:"string,min=2=E01"`
	Max      string `validate:"string,max=5=E03"`
	NonEmpty string `validate:"string,non_empty=true=E04"`
	Email    string `validate:"email,validemail=true=E05"`
}

var errs map[string]string

// InitError
func InitError() {
	errorMap := make(map[string]string, 0)
	errorMap["E01"] = "Min value int should be 1"
	errorMap["E02"] = "Max value should be 100"
	errorMap["E03"] = "Max value for string should be 10"
	errorMap["E04"] = "Cannot be empty"
	errorMap["E05"] = "Must be valid email format"
	errorMap["E06"] = "Greater than or equal error"
	errorMap["E07"] = "Less than or equal error"
	errorMap["E08"] = "Greater than error"
	errorMap["E09"] = "Less than error"
	errs = errorMap
}
func TestStringMinMaxEmail(t *testing.T) {
	InitError()
	NewError(errs)

	ss := StringStruct{
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

func TestValidStringMinMaxEmail(t *testing.T) {
	InitError()
	NewError(errs)

	ss := StringStruct{
		Max:      "lima",
		Min:      "dua",
		NonEmpty: "ada",
		Email:    "zainulmasadi90@gmail.com",
	}

	nuxErrors := ValidateStruct(ss)

	assert.Equal(t, 0, len(nuxErrors), "total of error should be 0")
}

type NumericStruct struct {
	Min int `validate:"number,min=1=E01"`
	Max int `validate:"number,max=3=E02"`
	Gte int `validate:"number,gte=3=E06"`
	Lte int `validate:"number,lte=8=E07"`
	Gt  int `validate:"number,gt=3=E08"`
	Lt  int `validate:"number,lt=8=E09"`
}

func TestNumericMinMaxGteLte(t *testing.T) {
	InitError()
	NewError(errs)

	in := NumericStruct{
		Min: 0,
		Max: 4,
		Gte: 3,
		Lte: 8,
		Gt:  4,
		Lt:  7,
	}

	nuxErrors := ValidateStruct(in)

	assert.Equal(t, 6, len(nuxErrors), "total of error should be 6")
	assert.Equal(t, "E01", nuxErrors[0].Code, "Error ordering shoud be in struct field order, 1st is E01")
	assert.Equal(t, "E02", nuxErrors[1].Code, "Error ordering shoud be in struct field order, 2nd is E02")
	assert.Equal(t, "E06", nuxErrors[2].Code, "Error ordering shoud be in struct field order, 3rd is E06")
	assert.Equal(t, "E07", nuxErrors[3].Code, "Error ordering shoud be in struct field order, 4th is E07")
	assert.Equal(t, "E08", nuxErrors[4].Code, "Error ordering shoud be in struct field order, 5th is E08")
	assert.Equal(t, "E09", nuxErrors[5].Code, "Error ordering shoud be in struct field order, 6th is E09")

}

func TestValidNumericMinMaxGteLte(t *testing.T) {
	InitError()
	NewError(errs)

	in := NumericStruct{
		Min: 2,
		Max: 2,
		Gte: 2,
		Lte: 9,
		Gt:  2,
		Lt:  9,
	}

	nuxErrors := ValidateStruct(in)
	assert.Equal(t, 0, len(nuxErrors), "total of error should be 0")

}
