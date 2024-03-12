package comparsion_test

import (
	"fmt"
	"testing"

	"github.com/babenow/validators/comparsion"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

func TestGt(t *testing.T) {
	testCases := []struct {
		desc  string
		value int
		num   int
		ok    bool
	}{
		{
			desc:  "INT_5",
			value: 7,
			num:   5,
			ok:    true,
		},
		{
			desc:  "Error INT_5",
			value: 3,
			num:   5,
			ok:    false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := validation.Validate(tC.value, validation.By(comparsion.Gt(tC.num, fmt.Sprintf("more %d", tC.num))))
			if tC.ok {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
