package helper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/validator/errors"
)

var errs map[string]string

func InitError() {
	errorMap := make(map[string]string, 0)
	errorMap["E01"] = "Min value int should be 1"
	errs = errorMap
}

func TestFindError(t *testing.T) {
	InitError()
	errors.InitializeError(errs)

	err := FindError("E01")
	errTwo := FindError("E02")

	assert.NotNil(t, err, "Error should be not nil")
	assert.Equal(t, "E01", err.Code, "Error code should be E01")

	assert.NotNil(t, errTwo, "Error should be not nil")
	assert.Equal(t, "XXXX", errTwo.Code, "Error code should be XXXX")
}

func TestParseBool(t *testing.T) {
	valBoolTrue := ParseBool("true")
	valBoolFalse := ParseBool("false")
	valBoolInvalid := ParseBool("bebas")

	assert.Equal(t, true, valBoolTrue, "Should be have true value")
	assert.Equal(t, false, valBoolFalse, "Should be have false value")
	assert.Equal(t, false, valBoolInvalid, "Should be have false value")
}

func TestParseInt(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Valid Number",
			args{
				num: "1",
			},
			1,
		},
		{
			"InValid Number",
			args{
				num: "abc",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInt(tt.args.num); got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitError(t *testing.T) {
	type args struct {
		erroStr string
	}
	tests := []struct {
		name string
		args args
		want *domain.NuxTag
	}{
		{
			"Valid tag for error code",
			args{
				erroStr: "gte=5=E012",
			},
			&domain.NuxTag{
				ErrorCode: "E012",
				Tag:       "gte",
				Value:     "5",
			},
		},
		{
			"InValid error code",
			args{
				erroStr: "gte=E012",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitError(tt.args.erroStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitError() = %v, want %v", got, tt.want)
			}
		})
	}
}
