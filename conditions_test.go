package validators_test

import (
	"testing"

	"github.com/babenow/validators"
	"github.com/stretchr/testify/assert"
)

func TestRequiredIF(t *testing.T) {
	// Test case 1: Condition is true, value is nil
	ruleFunc := validators.RequiredIF(true)
	if err := ruleFunc(nil); err == nil {
		t.Error("Expected error, got nil")
	}

	// Test case 2: Condition is true, value is not nil
	if err := ruleFunc("test"); err != nil {
		assert.Error(t, err, "Expected nil, got error")
	}

	// Test case 3: Condition is false
	ruleFunc = validators.RequiredIF(false)
	if err := ruleFunc(nil); err != nil {
		assert.Error(t, err, "Expected nil, got error")
	}
}
