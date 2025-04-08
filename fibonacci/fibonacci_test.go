package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	result1 := Fibonacci(1)
	result2 := Fibonacci(2)
	result3 := Fibonacci(3)
	result4 := Fibonacci(4)
	result5 := Fibonacci(5)

	assert.Equal(t, 0, result1)
	assert.Equal(t, 1, result2)
	assert.Equal(t, 1, result3)
	assert.Equal(t, 2, result4)
	assert.Equal(t, 3, result5)
}
