package validators_test

import (
	"fmt"
	"testing"

	"github.com/babenow/validators"
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
		{
			desc:  "Error INT_5 EQUAL",
			value: 5,
			num:   5,
			ok:    false,
		},
	}

	for _, tC := range testCases {
		t.Run(
			tC.desc, func(t *testing.T) {
				err := validation.Validate(
					tC.value, validation.By(validators.Gt(tC.num, fmt.Sprintf("more %d", tC.num))),
				)
				if tC.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}

func TestLt(t *testing.T) {
	testCases := []struct {
		desc  string
		value int
		num   int
		ok    bool
	}{
		{
			desc:  "INT_5",
			value: 5,
			num:   7,
			ok:    true,
		},
		{
			desc:  "Error INT_5",
			value: 7,
			num:   5,
			ok:    false,
		},
		{
			desc:  "Error INT_5 EQUAL",
			value: 5,
			num:   5,
			ok:    false,
		},
	}

	for _, tC := range testCases {
		t.Run(
			tC.desc, func(t *testing.T) {
				err := validation.Validate(
					tC.value, validation.By(validators.Lt(tC.num, fmt.Sprintf("more %d", tC.num))),
				)
				if tC.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}

func TestLte(t *testing.T) {
	testCases := []struct {
		desc  string
		value int
		num   int
		ok    bool
	}{
		{
			desc:  "INT_5",
			value: 5,
			num:   7,
			ok:    true,
		},
		{
			desc:  "Error INT_5",
			value: 7,
			num:   5,
			ok:    false,
		},
		{
			desc:  "Error INT_5 EQUAL",
			value: 5,
			num:   5,
			ok:    true,
		},
	}

	for _, tC := range testCases {
		t.Run(
			tC.desc, func(t *testing.T) {
				err := validation.Validate(
					tC.value, validation.By(validators.Lte(tC.num, fmt.Sprintf("more %d", tC.num))),
				)
				if tC.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}

func TestGte(t *testing.T) {
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
			value: 5,
			num:   7,
			ok:    false,
		},
		{
			desc:  "Error INT_5 EQUAL",
			value: 5,
			num:   5,
			ok:    true,
		},
	}

	for _, tC := range testCases {
		t.Run(
			tC.desc, func(t *testing.T) {
				err := validation.Validate(
					tC.value, validation.By(validators.Gte(tC.num, fmt.Sprintf("more %d", tC.num))),
				)
				if tC.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}

func TestBetween(t *testing.T) {
	testCases := []struct {
		desc  string
		value int
		left  int
		right int
		ok    bool
	}{
		{
			desc:  "Between 5-10 (7)",
			value: 7,
			left:  5,
			right: 10,
			ok:    true,
		},
		{
			desc:  "Not Between 5-10 (lt)",
			value: 3,
			left:  5,
			right: 10,
			ok:    false,
		},
		{
			desc:  "Not Between 5-10 (gt)",
			value: 11,
			left:  5,
			right: 10,
			ok:    false,
		},
		{
			desc:  "Between 5-10 (left)",
			value: 5,
			left:  5,
			right: 10,
			ok:    true,
		},
		{
			desc:  "Between 5-10 (right)",
			value: 10,
			left:  5,
			right: 10,
			ok:    true,
		},
	}
	for _, tC := range testCases {

		t.Run(
			tC.desc, func(t *testing.T) {
				err := validation.Validate(
					tC.value, validation.By(
						validators.Between(
							tC.left, tC.right, fmt.Sprintf("between %d-%d", tC.left, tC.right),
						),
					),
				)
				if tC.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}

func TestNot(t *testing.T) {
	testcases := []struct {
		desc  string
		num   int
		value int
		ok    bool
	}{
		{
			desc:  "OK",
			num:   3,
			value: 5,
			ok:    true,
		},
	}

	for _, tc := range testcases {
		t.Run(
			tc.desc, func(t *testing.T) {
				err := validation.Validate(tc.value, validation.By(validators.Not(tc.num, "error")))
				if tc.ok {
					assert.Nil(t, err)
				} else {
					assert.Error(t, err)
				}
			},
		)
	}
}
