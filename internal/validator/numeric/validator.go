package numeric

import (
	"github.com/zainul/nux/internal/domain"
	"github.com/zainul/nux/internal/helper"
	"github.com/zainul/nux/internal/validator/usecase"
)

const (
	// NonZero tag
	NonZero = "non_zero"
	// Min tag
	Min = "min"
	// Max tag
	Max = "max"
	// GreaterThanEqual tag
	GreaterThanEqual = "gte" // Same as 'min' above example gte=3 so min value should be 1,2,3
	// LessThanEqual tag
	LessThanEqual = "lte" // Same as 'max' above example lte=10 so max value should be 10,9,8,7
	// GreaterThan  tag
	GreaterThan = "gt"
	// LessThan tag
	LessThan = "lt"
)

// New ...
func New() usecase.Validator {
	errMap := make(map[string]string, 0)
	return &validator{
		ErrorCode: errMap,
	}
}

// Validator ...
type validator struct {
	ErrorCode        map[string]string
	Min              *int
	Max              *int
	GreaterThanEqual *int
	LessThanEqual    *int
	GreaterThan      *int
	LessThan         *int
	NonZero          bool
}

// Validate implementation of numeric validation
func (v *validator) Validate(val interface{}) (bool, *domain.NuxError) {
	if l, ok := val.(int); ok {
		return v.validateInt(l)
	}

	if l, ok := val.(int64); ok {
		return v.validateInt64(l)
	}

	if l, ok := val.(float64); ok {
		return v.validateFloat64(l)
	}

	if l, ok := val.(float32); ok {
		return v.validateFloat32(l)
	}

	return true, nil
}

func (v *validator) validateInt(l int) (bool, *domain.NuxError) {
	if v.NonZero && l == 0 {
		// fmt.Println("non zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[NonZero])
	}

	if v.Min != nil && l < *v.Min {
		// fmt.Println("min detected", l, v)
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max != nil && l > *v.Max {
		// fmt.Println("max zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[Max])
	}

	if v.LessThanEqual != nil && l <= *v.LessThanEqual {
		// fmt.Println("lte detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThanEqual])
	}

	if v.GreaterThanEqual != nil && l >= *v.GreaterThanEqual {
		// fmt.Println("gte detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThanEqual])
	}

	if v.LessThan != nil && l < *v.LessThan {
		// fmt.Println("lt detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThan])
	}

	if v.GreaterThan != nil && l > *v.GreaterThan {
		// fmt.Println("gt detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThan])
	}
	return true, nil
}

func (v *validator) validateInt64(l int64) (bool, *domain.NuxError) {
	if v.NonZero && l == 0 {
		// fmt.Println("non zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[NonZero])
	}

	if v.Min != nil && l < int64(*v.Min) {
		// fmt.Println("min detected", l, v)
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max != nil && l > int64(*v.Max) {
		// fmt.Println("max zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[Max])
	}

	if v.LessThanEqual != nil && l <= int64(*v.LessThanEqual) {
		// fmt.Println("lte detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThanEqual])
	}

	if v.GreaterThanEqual != nil && l >= int64(*v.GreaterThanEqual) {
		// fmt.Println("gte detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThanEqual])
	}

	if v.LessThan != nil && l < int64(*v.LessThan) {
		// fmt.Println("lt detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThan])
	}

	if v.GreaterThan != nil && l > int64(*v.GreaterThan) {
		// fmt.Println("gt detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThan])
	}
	return true, nil
}

func (v *validator) validateFloat64(l float64) (bool, *domain.NuxError) {
	if v.NonZero && l == 0 {
		// fmt.Println("non zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[NonZero])
	}

	if v.Min != nil && l < float64(*v.Min) {
		// fmt.Println("min detected", l, v)
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max != nil && l > float64(*v.Max) {
		// fmt.Println("max zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[Max])
	}

	if v.LessThanEqual != nil && l <= float64(*v.LessThanEqual) {
		// fmt.Println("lte detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThanEqual])
	}

	if v.GreaterThanEqual != nil && l >= float64(*v.GreaterThanEqual) {
		// fmt.Println("gte detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThanEqual])
	}

	if v.LessThan != nil && l < float64(*v.LessThan) {
		// fmt.Println("lt detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThan])
	}

	if v.GreaterThan != nil && l > float64(*v.GreaterThan) {
		// fmt.Println("gt detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThan])
	}
	return true, nil
}

func (v *validator) validateFloat32(l float32) (bool, *domain.NuxError) {
	if v.NonZero && l == 0 {
		// fmt.Println("non zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[NonZero])
	}

	if v.Min != nil && l < float32(*v.Min) {
		// fmt.Println("min detected", l, v)
		return false, helper.FindError(v.ErrorCode[Min])
	}

	if v.Max != nil && l > float32(*v.Max) {
		// fmt.Println("max zero detected", l, v)
		return false, helper.FindError(v.ErrorCode[Max])
	}

	if v.LessThanEqual != nil && l <= float32(*v.LessThanEqual) {
		// fmt.Println("lte detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThanEqual])
	}

	if v.GreaterThanEqual != nil && l >= float32(*v.GreaterThanEqual) {
		// fmt.Println("gte detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThanEqual])
	}

	if v.LessThan != nil && l < float32(*v.LessThan) {
		// fmt.Println("lt detected", l, v)
		return false, helper.FindError(v.ErrorCode[LessThan])
	}

	if v.GreaterThan != nil && l > float32(*v.GreaterThan) {
		// fmt.Println("gt detected", l, v)
		return false, helper.FindError(v.ErrorCode[GreaterThan])
	}
	return true, nil
}

// SetValueFromTag set value base on struct tag
func (v *validator) SetValueFromTag(field *domain.NuxTag) {
	if field != nil {

		switch field.Tag {
		case Min:
			min := helper.ParseInt(field.Value)
			v.Min = &min
		case Max:
			max := helper.ParseInt(field.Value)
			v.Max = &max
		case GreaterThanEqual:
			gte := helper.ParseInt(field.Value)
			v.GreaterThanEqual = &gte
		case LessThanEqual:
			lte := helper.ParseInt(field.Value)
			v.LessThanEqual = &lte
		case GreaterThan:
			gt := helper.ParseInt(field.Value)
			v.GreaterThan = &gt
		case LessThan:
			lt := helper.ParseInt(field.Value)
			v.LessThan = &lt
		case NonZero:
			v.NonZero = helper.ParseBool(field.Value)
		}

		v.ErrorCode[field.Tag] = field.ErrorCode
	}
}
