package main
//will have access to the things available in the main package

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T){
	assert.Equal(t, 123, 123, "they should be equal")
}