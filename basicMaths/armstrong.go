package basicmaths

import "math"

// Check a number is armstrong number or not
// time complexity : BigO(logN)
func Armstrong(n int) bool {
	digits := CountDigitFaster(n)
	sum := 0
	num := n

	for num > 0 {
		eachDigit := num % 10
		sum += int(math.Pow(float64(eachDigit), float64(digits)))
		num /= 10
	}

	if sum == n {
		return true
	}
	return false
}
