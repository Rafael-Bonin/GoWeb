package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideSucces(t *testing.T) {
	// validate valid values
	n1 := 10.0
	n2 := 8.0
	expectedResult := 1.25

	result, _ := Divide(n1, n2)
	
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivideError(t *testing.T) {
	// validate invalid values
	n1 := 10.0
	n2 := 0.0
	_, err := Divide(n1, n2)
	
	assert.NotNil(t, err)
}