package domain

// NuxError is struct for result validation
type NuxError struct {
	Code    string
	Message string
}

// NuxTag is struct for set value from struct tag
type NuxTag struct {
	Tag       string
	Value     string
	ErrorCode string
}
