package validators

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Number ...
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Gt GtreaterThe...
func Gt[T Number](num T, msg string) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(T) < num || value.(T) == num {
			return fmt.Errorf("%s", msg)
		}

		return nil
	}
}

// Lt LessThe....
func Lt[T Number](num T, msg string) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(T) > num || value.(T) == num {
			return fmt.Errorf("%s", msg)
		}

		return nil
	}
}
