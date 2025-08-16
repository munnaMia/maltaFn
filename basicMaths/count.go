package basicmaths

import "math"

// take n as input and tells how many digits on the n
// time complexity : Big O(logN)
func CountDigit(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n / 10
	}
	return count
}

// take n as input and tells how many digits on the n
// time complexity : Big O(1)
func CountDigitFaster(n int) int {
	return int(math.Log10(float64(n)) + 1)
}
