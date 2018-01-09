package fields

import (
	"strconv"

	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/validators"
)

// IntegerErrors is a code-error mapping for Integer field.
var IntegerErrors = map[uint]string{
	codes.Required: "This field is required.",
	codes.Invalid:  "Ensure this value is valid integer.",
	codes.MinValue: "Ensure this value is greater than or equal to %d.",
	codes.MaxValue: "Ensure this value is less than or equal to %d.",
}

// Integer is integer field.
type Integer struct {
	Name       string
	Validators []validators.IntegerValidator
	Required   bool
	Default    int
	Errors     map[uint]string
	ErrorFunc  ErrorFunc

	AllowStrings bool
}

// IsRequired returns true if field is required.
func (field *Integer) IsRequired() bool {
	return field.Required
}

// GetDefault returns the default value.
func (field *Integer) GetDefault() interface{} {
	return field.Default
}

// GetName returns field name.
func (field *Integer) GetName() string {
	return field.Name
}

// GetError returns error by code.
func (field *Integer) GetError(code uint, value interface{}, parameters ...interface{}) error {
	return getError(field, code, value, field.Errors, IntegerErrors, field.ErrorFunc, parameters...)
}

// Validate check and clean an input value.
func (field *Integer) Validate(v interface{}) (interface{}, error) {
	var value int

	switch v := v.(type) {
	case int:
		value = v
	case float64:
		value = int(v)
	case string:
		if field.AllowStrings {
			var err error

			value, err = strconv.Atoi(v)

			if err != nil {
				return nil, field.GetError(codes.Invalid, v)
			}
		} else {
			return nil, field.GetError(codes.Invalid, v)
		}
	default:
		return nil, field.GetError(codes.Invalid, v)
	}

	for _, validator := range field.Validators {
		var err *validators.Error

		value, err = validator.Validate(value)

		if err != nil {
			return nil, field.GetError(err.Code, v, err.Parameters...)
		}
	}

	return value, nil
}
