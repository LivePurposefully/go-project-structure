package maths

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubtractPositiveIntegers(t *testing.T) {
	assert := assert.New(t)
	result := Sub(2,3)
	assert.Equal(-1, result, "2 - 3 should be -1")
}