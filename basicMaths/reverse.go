package basicmaths

// take input n as int and return reverse of n
// time complexity : Big O(logN)
func ReverseDigit(n int) int {
	reverse_result := 0
	for n > 0 {
		reverse_result = reverse_result*10 + n%10
		n /= 10
	}
	return reverse_result
}
