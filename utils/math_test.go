package utils

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{-1, -5, -1},
		{0, 0, 0},
		{100, 50, 100},
		{-10, 10, 10},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 1},
		{5, 3, 3},
		{-1, -5, -5},
		{0, 0, 0},
		{100, 50, 50},
		{-10, 10, -10},
	}

	for _, test := range tests {
		result := Min(test.a, test.b)
		if result != test.expected {
			t.Errorf("Min(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{123, 123},
		{-123, 123},
		{1, 1},
		{-1, 1},
	}

	for _, test := range tests {
		result := Abs(test.input)
		if result != test.expected {
			t.Errorf("Abs(%d) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestAbsFloat(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.5, 5.5},
		{-5.5, 5.5},
		{0.0, 0.0},
		{123.456, 123.456},
		{-123.456, 123.456},
	}

	for _, test := range tests {
		result := AbsFloat(test.input)
		if result != test.expected {
			t.Errorf("AbsFloat(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{17, true},
		{25, false},
		{29, true},
		{100, false},
		{101, true},
	}

	for _, test := range tests {
		result := IsPrime(test.input)
		if result != test.expected {
			t.Errorf("IsPrime(%d) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 8, 4},
		{15, 10, 5},
		{7, 3, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{-12, 8, 4},
		{12, -8, 4},
		{-12, -8, 4},
		{48, 18, 6},
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		if result != test.expected {
			t.Errorf("GCD(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 8, 24},
		{15, 10, 30},
		{7, 3, 21},
		{0, 5, 0},
		{5, 0, 0},
		{0, 0, 0},
		{4, 6, 12},
		{-4, 6, 12},
		{4, -6, 12},
		{-4, -6, 12},
	}

	for _, test := range tests {
		result := LCM(test.a, test.b)
		if result != test.expected {
			t.Errorf("LCM(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{-1, 0},
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		base, exp int
		expected  int
	}{
		{2, 0, 1},
		{2, 1, 2},
		{2, 3, 8},
		{5, 2, 25},
		{10, 3, 1000},
		{-2, 2, 4},
		{-2, 3, -8},
		{3, -1, 0},
		{0, 5, 0},
		{1, 100, 1},
	}

	for _, test := range tests {
		result := Pow(test.base, test.exp)
		if result != test.expected {
			t.Errorf("Pow(%d, %d) = %d, expected %d", test.base, test.exp, result, test.expected)
		}
	}
}