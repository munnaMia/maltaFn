package basicmaths

// Check a number is prime or not.
// time complexity : Big O(N)
func Prime(n int) bool {
	counter := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			counter++
		}
	}

	// why 2 because a prime number has only 2 divisor 1 and itself.
	if counter == 2 {
		return true
	}
	return false
}

// Check a number is prime or not.
// time complexity : Big O(sqrt(N))
func PrimeFaster(n int)bool {
	divisorsOfN := DivisorsFaster(n)

	if len(divisorsOfN) == 2 {
		return true
	}
	return false
}
