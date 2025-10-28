package arrays

// Find the Majority Element that occurs more than N/2 times
// time complexity Big O(n)
// space complexity Big O(1)

func FindMajority(nums []int) int {
	count := 0
	elem := 0

	for _, v := range nums {
		if count == 0 {
			elem = v
		}

		if elem == v {
			count++
		} else {
			count--
		}

	}

	cnt1 := 0
	for _, v := range nums {
		if v == elem {
			cnt1 += 1
		}
	}

	if cnt1 > len(nums)/2 {
		return elem
	}

	return -1

}
