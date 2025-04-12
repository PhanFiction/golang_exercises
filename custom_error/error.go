package customerror

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, &err
	}
	return x * x, nil
}
