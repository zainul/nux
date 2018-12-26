package helper

import (
	"strconv"
	"strings"

	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/validator/errors"
)

// FindError to find error message in map
func FindError(errCode string) *domain.NuxError {
	if val, ok := errors.ErrorMap[errCode]; ok {
		return &domain.NuxError{
			Code:    errCode,
			Message: val,
		}
	}
	return &domain.NuxError{
		Code:    "XXXX",
		Message: "Unexpected error",
	}
}

// SplitError ....
func SplitError(erroStr string) *domain.NuxTag {
	splitted := strings.Split(erroStr, "=")

	if len(splitted) == 3 {
		return &domain.NuxTag{
			Tag:       splitted[0],
			Value:     splitted[1],
			ErrorCode: splitted[2],
		}
	}
	return nil
}

// ParseInt ...
func ParseInt(num string) int {
	if val, err := strconv.Atoi(num); err != nil {
		return 0
	} else {
		return val
	}
}

// ParseBool  ....
func ParseBool(bol string) bool {
	if val, err := strconv.ParseBool(bol); err != nil {
		return false
	} else {
		return val
	}
}
