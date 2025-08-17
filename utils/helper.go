package utils

// return the biggest number form n1 and n2 and if n1 == n2 then n1 will be returned.
func Max[T int | float64](n1, n2 T) T {
	if n1 >= n2 {
		return n1
	} else {
		return n2
	}
}

// return the smallest number form n1 and n2 and if n1 == n2 then n1 will be returned.
func Min[T int | float64](n1, n2 T) T {
	if n1 <= n2 {
		return n1
	} else {
		return n2
	}
}

