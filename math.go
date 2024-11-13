package gutil

import (
	"errors"
	"math"
)

var ClampErr = errors.New("clamp: min > max")

func Clamp(number float64, min float64, max float64) (float64, error) {
	if min > max {
		return 0, ClampErr
	}
	number = math.Min(max, number)
	number = math.Max(min, number)
	return number, nil
}
