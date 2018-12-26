package usecase

import "github.com/zainul/nux/internal/domain"

// Validator ...
type Validator interface {
	Validate(val interface{}) (bool, *domain.NuxError)
	SetValueFromTag(field *domain.NuxTag)
}
