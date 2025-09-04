package sorting

// Bubble sort is a simple sorting algorithm that repeatedly steps through a list, compares adjacent elements, and swaps them if they are in the wrong order. This process continues until no more swaps are needed, indicating the list is sorted. Each pass through the list moves the largest unsorted element to its correct position at the end of the list, causing it to "bubble" up.

// Steps of Bubble Sort
// 		Start at the Beginning: Begin at the first element of the unsorted list.

// 		Compare Adjacent Elements: Compare the current element with the element next to it.

// 		Swap if Necessary: If the elements are in the wrong order (e.g., the current element is greater than the next for ascending order), swap them.

// 		Move to the Next Pair: Advance to the next pair of adjacent elements and repeat steps 2 and 3.

// 		Complete a Pass: Continue this process until you reach the end of the list. After the first pass, the largest element will be in its correct, final position at the end of the list.

// 		Repeat Passes: Start again from the beginning of the list to perform another pass.

// 		Reduce the Range: In each subsequent pass, you can reduce the number of comparisons by one, as the largest elements are already sorted and at the end of the list.

//		Check for Completion: The algorithm stops when a complete pass is made without any swaps occurring, as this indicates the list is fully sorted.

func BubbleSort(nums []int) []int {
	tempArr := nums
	for i := 0; i < len(tempArr)-1; i++ {
		for j := 0; j < len(tempArr)-i-1; j++ {
			if tempArr[j] > tempArr[j+1] {
				tempArr[j], tempArr[j+1] = tempArr[j+1], tempArr[j]
			}
		}
	}
	return tempArr
}
