package nux

import (
	"github.com/zainul/nux/internal"
	"github.com/zainul/nux/internal/domain"
)

// NewError ...
func NewError(errs map[string]string) {
	internal.NewError(errs)
}

// ValidateStruct  ....
func ValidateStruct(strct interface{}) []domain.NuxError {
	return internal.ValidateStruct(strct)
}
