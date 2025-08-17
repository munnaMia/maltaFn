package basicmaths

import "github.com/munnaMia/maltaFn/utils"

// find the greatest common divisor form two inputs
// time complexity : Big O(min(n1, n2)) ----> Big O(N)
func GCD(n1, n2 int) int {
	gcd := 0
	min := utils.Min(n1, n2)

	for i := 1; i <= min; i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}
	return gcd
}

// find the greatest common divisor form two inputs
// faster based on itteration
// time complexity : Big O(min(n1, n2)) ----> Big O(N)
func GCDfast(n1, n2 int) int {
	for i := utils.Min(n1, n2); i > 1; i-- {
		if n1%i == 0 && n2%i == 0 {
			return i
		}
	}
	return 1
}

// find the greatest common divisor form two inputs using recursion and The Euclidean Algorithm
// time complexity : Big O(min(n1, n2)) ----> Big O(N)
func GCDfast1(n1, n2 int) int {

	if n1 > n2 {
		return GCDfast1(n1-n2, n2)
	}
	if n1 < n2 {
		return GCDfast1(n2-n1, n1)
	}

	if n1 == 0 {
		return n2
	} else {
		return n1
	}
}

// find the greatest common divisor form two inputs using loop and The Euclidean Algorithm
// time complexity : Big O(min(n1, n2)) ----> Big O(N)
func GCDfast2(n1, n2 int) int {
	for n1 > 0 && n2 > 0 {
		if n1 > n2 {
			n1 = n1 - n2
		} else {
			n2 = n2 - n1
		}
	}
	if n1 == 0 {
		return n2
	}
	return n1
}
