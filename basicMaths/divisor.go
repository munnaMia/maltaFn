package basicmaths

import "math"

// Take a input n and return its all divisors.
// time complexity : Big O(n).
func Divisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

// Take a input n and return its all divisors.
// time complexity : Big O(sqrt(N)).
func DivisorsFaster(n int) []int {
	sqrN := int(math.Sqrt(float64(n)))

	divisors := []int{}

	for i := 1; i <= sqrN; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
