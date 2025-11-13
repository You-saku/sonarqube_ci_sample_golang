package utils

import "math"

// Max は2つの整数の最大値を返します
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min は2つの整数の最小値を返します
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs は整数の絶対値を返します
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AbsFloat は浮動小数点数の絶対値を返します
func AbsFloat(x float64) float64 {
	return math.Abs(x)
}

// IsPrime は整数が素数かどうかを判定します
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// GCD は2つの整数の最大公約数を返します
func GCD(a, b int) int {
	a = Abs(a)
	b = Abs(b)
	
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM は2つの整数の最小公倍数を返します
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / GCD(a, b)
}

// Factorial は非負整数の階乗を返します
func Factorial(n int) int {
	if n < 0 {
		return 0 // 負数の場合は0を返す
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Pow は整数の累乗を返します
func Pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp < 0 {
		return 0 // 整数の負の累乗は0とする
	}
	
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}