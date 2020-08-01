package calculator

import "errors"

// Sum sum a by b
func Sum(a, b float64) float64 {
	return a + b
}

// Sub subtract a by b
func Sub(a, b float64) float64 {
	return a - b
}

// Mul multiply a by b
func Mul(a, b float64) float64 {
	return a * b
}

// Div divide a by b
func Div(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("division by zero")
	}
	return a / b, nil
}
