package errors

import (
	"sync"
)

// ErrorMap ...
var ErrorMap map[string]string
var once sync.Once

// InitializeError ...
func InitializeError(errs map[string]string) {

	once.Do(func() {
		ErrorMap = errs
	})
}
