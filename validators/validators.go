package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var AlphaSpace validator.Func = func(fieldLevel validator.FieldLevel) bool {
	value, ok := fieldLevel.Field().Interface().(string)
	if ok {
		valid, err := regexp.Match(`(?i)^[a-z][a-z ]*[a-z]$`, []byte(value))
		if err != nil {
			return false
		}
		if valid {
			return true
		}
	}
	return false
}
var AlphaNumSpace validator.Func = func(fieldLevel validator.FieldLevel) bool {
	value, ok := fieldLevel.Field().Interface().(string)
	if ok {
		valid, err := regexp.Match(`(?i)^[a-z][a-z0-9 ]*[a-z0-9]$`, []byte(value))
		if err != nil {
			return false
		}
		if valid {
			return true
		}
	}
	return false
}
var Username validator.Func = func(fieldLevel validator.FieldLevel) bool {
	value, ok := fieldLevel.Field().Interface().(string)
	if ok {
		valid, err := regexp.Match(`(?i)^[a-z][a-z0-9_.]{6,18}[a-z0-9]$`, []byte(value))
		if err != nil {
			return false
		}
		if valid {
			return true
		}
	}
	return false
}
var Password validator.Func = func(fieldLevel validator.FieldLevel) bool {
	value, ok := fieldLevel.Field().Interface().(string)
	if ok {
		re := []string{`[a-z]`, `[A-Z]`, `[0-9]`, `[\W]`}
		var valid bool
		var err error
		for _, exp := range re {
			valid, err = regexp.Match(exp, []byte(value))
			if err != nil {
				return false
			} else if !valid {
				return false
			}
		}
		return true
	}
	return false
}
