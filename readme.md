# Nux

![Alt text](logo.png?raw=true "Clean Architecture")

[![Go Report Card](https://goreportcard.com/badge/github.com/zainul/nux)](https://goreportcard.com/report/github.com/zainul/nux) [![CircleCI](https://circleci.com/gh/zainul/nux.svg?style=svg)](https://circleci.com/gh/zainul/nux)
[![codecov](https://codecov.io/gh/zainul/nux/branch/master/graph/badge.svg)](https://codecov.io/gh/zainul/nux)



nux is golang validator that allow to setting the error message dynamically base on error code.
currently support for:

- text validation
- number validation
- email validation

## Install

```
go get github.com/zainul/nux
```

## Usage

in definition of struct

```
type StringStruct struct {
	Min      string `validate:"string,min=2=E01"`
	Max      string `validate:"string,max=5=E03"`
	NonEmpty string `validate:"string,non_empty=true=E04"`
	Email    string `validate:"email,validemail=true=E05"`
}

type NumericStruct struct {
	Min int `validate:"number,min=1=E01"`
	Max int `validate:"number,max=3=E02"`
	Gte int `validate:"number,gte=3=E06"`
	Lte int `validate:"number,lte=8=E07"`
	Gt  int `validate:"number,gt=3=E08"`
	Lt  int `validate:"number,lt=8=E09"`
}

```


when validate in the code

```
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

func main() {
    InitError()

	NewError(errs)

	ss := StringStruct{
		Max:      "ulalala",
		Email:    "email",
		Min:      "u",
		NonEmpty: "",
	}
	
    listOfError := ValidateStruct(ss)

    for _, err := range listOfError {
        fmt.Println(err.Code, err.Message)
    }
}
```
