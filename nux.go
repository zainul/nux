package nux

import (
	"github.com/zainul/nux/internal"
	"github.com/zainul/nux/internal/domain"
)

// NewError initiate the list of error
func NewError(errs map[string]string) {
	internal.NewError(errs)
}

// ValidateStruct  is function to use validate the struct
func ValidateStruct(strct interface{}) []domain.NuxError {
	return internal.ValidateStruct(strct)
}
