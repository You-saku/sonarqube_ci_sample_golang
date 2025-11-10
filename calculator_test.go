package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed numbers", 5.0, -3.0, 2.0},
		{"zero addition", 0.0, 5.0, 5.0},
		{"decimal numbers", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"zero subtraction", 5.0, 0.0, 5.0},
		{"decimal numbers", 5.5, 2.3, 3.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5.0, 3.0, 15.0},
		{"negative numbers", -5.0, -3.0, 15.0},
		{"mixed numbers", 5.0, -3.0, -15.0},
		{"zero multiplication", 0.0, 5.0, 0.0},
		{"decimal numbers", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", 10.0, 2.0, 5.0, false},
		{"negative numbers", -10.0, -2.0, 5.0, false},
		{"mixed numbers", 10.0, -2.0, -5.0, false},
		{"decimal result", 7.0, 2.0, 3.5, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name           string
		base, exponent float64
		want           float64
	}{
		{"positive base and exponent", 2.0, 3.0, 8.0},
		{"base to power of 0", 5.0, 0.0, 1.0},
		{"base to power of 1", 5.0, 1.0, 5.0},
		{"zero base", 0.0, 3.0, 0.0},
		{"negative base", -2.0, 2.0, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.base, tt.exponent); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Power(%v, %v) = %v, want %v", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"positive even", 4, true},
		{"positive odd", 5, false},
		{"zero", 0, true},
		{"negative even", -4, true},
		{"negative odd", -5, false},
		{"large even", 1000, true},
		{"large odd", 1001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Errorf("IsEven(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"a greater than b", 10.0, 5.0, 10.0},
		{"b greater than a", 5.0, 10.0, 10.0},
		{"equal numbers", 5.0, 5.0, 5.0},
		{"negative numbers", -10.0, -5.0, -5.0},
		{"mixed numbers", -5.0, 10.0, 10.0},
		{"decimal numbers", 3.14, 2.71, 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.a, tt.b); got != tt.want {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"a less than b", 5.0, 10.0, 5.0},
		{"b less than a", 10.0, 5.0, 5.0},
		{"equal numbers", 5.0, 5.0, 5.0},
		{"negative numbers", -10.0, -5.0, -10.0},
		{"mixed numbers", -5.0, 10.0, -5.0},
		{"decimal numbers", 2.71, 3.14, 2.71},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.a, tt.b); got != tt.want {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10.0, 3.0)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(10.0, 3.0)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(10.0, 3.0)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2.0, 10.0)
	}
}