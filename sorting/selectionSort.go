package sorting

/*
Selection Sort Algorithm:
--------------------------
Iterate through the array: Start with the first element and traverse the array
up to the second-to-last element. This outer loop will mark the boundary between
the sorted and unsorted parts of the array. Let's call the index of the current
element i.

Find the minimum element: For each iteration of the outer loop (from index i to
the end of the array), find the minimum element in the unsorted subarray. Keep
track of the index of this minimum element.

Swap: Once the inner loop completes and the minimum element's index is found,
swap the minimum element with the element at the current position i. This effectively
places the smallest element in its correct sorted position.

Repeat: Increment i and repeat the process. The outer loop continues until the entire
array is sorted.
*/

func SelectionSort(n []int) []int {
	length := len(n)
	for i := 0; i < length-1; i++ {
		minValueIdx := i
		for j := i + 1; j < length; j++ {
			if n[minValueIdx] > n[j] {
				minValueIdx = j // select the min value index
			}
		}
		if n[i] > n[minValueIdx] {
			n[i], n[minValueIdx] = n[minValueIdx], n[i]
		}
	}
	return n
}
