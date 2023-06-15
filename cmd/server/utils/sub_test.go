package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	n1 := 10
	n2 := 5
	expect := 5
	result := Sub(n1, n2)
	
	assert.Equal(t, expect, result, "deve subtrair o segundo numero do primeiro")
}