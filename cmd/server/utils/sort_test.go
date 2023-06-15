package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	var list []int = []int{10,9,8,7,6,5,4,3,2,1,0}
	var expectedResult []int = []int{0,1,2,3,4,5,6,7,8,9,10}

	result := Sort(list)

	assert.Equal(t, expectedResult, result, "deve, ser iguais")
}