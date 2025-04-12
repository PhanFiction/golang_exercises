package customerror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	res, err := Sqrt(-2)
	assert.Equal(t, "cannot Sqrt negative number: -2", err.Error())
	assert.Equal(t, float64(-2), res)

	res2, err2 := Sqrt(2)
	assert.Equal(t, res2, float64(4))
	assert.Equal(t, err2, nil)
}
