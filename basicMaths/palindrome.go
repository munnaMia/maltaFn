package basicmaths

// Check a number is palindrome or not.
// time complexity : Big O(logN)
func Palindrome(n int) bool {
	reverse_n := ReverseDigit(n)
	if n == reverse_n {
		return true
	}
	return false
}
