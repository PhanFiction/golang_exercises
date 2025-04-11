package customstringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringer(t *testing.T) {
	input := IPAddr{192, 168, 1, 1}
	expected := "192.168.1.1"
	result := input.String()
	assert.Equal(t, expected, result)

	input2 := IPAddr{8, 8, 8, 8}
	expected2 := "8.8.8.8"
	result2 := input2.String()
	assert.Equal(t, expected2, result2)
}
