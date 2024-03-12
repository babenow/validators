package validators_test

import (
	"testing"

	"github.com/babenow/validators"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

func TestRerquiredIF(t *testing.T) {
	testCases := []struct {
		desc  string
		cond  bool
		value interface{}
		ok    bool
	}{
		{desc: "TRUE OK", cond: true, value: "required", ok: true},
		{desc: "TRUE NOT OK", cond: true, value: nil, ok: false},
		{desc: "FALSE OK", cond: false, value: nil, ok: true},
		{desc: "FALSE NOT OK", cond: false, value: "required", ok: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := validation.Validate(tC.value, validation.By(validators.RequiredIF(tC.cond)))

			if tC.ok {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
