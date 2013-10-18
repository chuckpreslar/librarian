package librarian

import (
	"fmt"
)

type Attribute struct {
	name       string
	typ        Type
	validators []Validator
}

func (a *Attribute) ValidatesWith(validator Validator) *Attribute {
	a.validators = append(a.validators, validator)
	return a
}

func (a *Attribute) ConvertValue(value interface{}) interface{} {
	switch a.typ {
	case Integer:
		return a.ConvertInteger(value)
	case String:
		return a.ConvertString(value)
	default:
		panic(fmt.Errorf("unknown type, was %T", value))
	}
}

func (a *Attribute) ConvertInteger(value interface{}) int64 {
	switch value.(type) {
	case int:
		return int64(value.(int))
	case int8:
		return int64(value.(int8))
	case int16:
		return int64(value.(int16))
	case int32:
		return int64(value.(int32))
	case int64:
		return int64(value.(int64))
	default:
		panic(fmt.Errorf("value was not a valid int type, was %T", value))
	}
}

func (a *Attribute) ConvertString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}
