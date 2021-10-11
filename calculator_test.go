package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestManyAdd(t *testing.T) {
	t.Parallel()
	//	var want float64 = 4
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2, 2, 2}, want: 6},
		{inputs: []float64{5, 0, 5, 5}, want: 15},
	}
	for _, tc := range testCases {
		got := calculator.AddMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("AddMany(%f): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestSubstractMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{3, 4, 5}, want: -6},
		{inputs: []float64{10, 1, 2, 3}, want: 4},
	}
	for _, tc := range testCases {
		got := calculator.SubstractMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("SubstractMany %f): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestMutiplyMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2, 2, 2}, want: 8},
		{inputs: []float64{10, -5, 2}, want: -100},
	}

	for _, tc := range testCases {
		got := calculator.MultiplyMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("Multiply %f): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestDivideMany(t *testing.T) {
	t.Parallel()
	//	var want float64 = 4
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{4, 2}, want: 2},
		{inputs: []float64{10, 2, 5}, want: 1},
		{inputs: []float64{1000, 10, 10}, want: 10},
	}
	for _, tc := range testCases {
		got, err := calculator.DivideMany(tc.inputs...)
		if err != nil {
			t.Fatalf("Divide (%f) want no error for valid input, got %v", tc.inputs, err)
		}
		if tc.want != got {
			t.Errorf("Divide %f): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.DivideMany(10, 5, 0, 7)
	if err == nil {
		t.Error("Divide(10, 5, 0, 7): want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 9, want: 3},
		{a: 3, want: 1.73},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("Sqrt (%f) want no error for valid input, got %v", tc.a, err)
		}
		if !closeEnough(tc.want, got, 0.01) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	// We use _ because we don't need the a value here
	_, err := calculator.Sqrt(-2)
	if err == nil {
		t.Error("Sqrt(-2): want error for invalid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
