package calculator

import (
	"errors"
	"math"
)

func AddMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	var result float64 = 0
	for _, value := range inputs {
		result += value
	}
	return result
}

func SubstractMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, value := range inputs[1:] {
		result -= value
	}
	return result
}

func MultiplyMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, value := range inputs[1:] {
		result *= value
	}
	return result
}

func DivideMany(inputs ...float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	result := inputs[0]
	for _, value := range inputs[1:] {
		if value == 0 {
			return 0, errors.New("division by zero not allowed")
		}
		result /= value
	}
	return result, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number not allowed")
	}
	return math.Sqrt(a), nil
}
