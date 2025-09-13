package sorting

/*
Quick sort is a popular and efficient sorting algorithm that uses a divide-and-conquer strategy. It works by picking an element from the array, called a pivot, and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then recursively sorted. The process repeats until the entire array is sorted.

Steps of Quick Sort
Select a Pivot: Choose an element from the array to be the pivot. The choice of pivot can greatly affect the algorithm's performance. Common strategies include picking the first element, the last element, a random element, or the median of three elements.

Partitioning: Rearrange the array so that all elements smaller than the pivot come before it, and all elements larger than the pivot come after it. Elements equal to the pivot can go on either side. After this step, the pivot is in its final sorted position.

Recursive Sort: Recursively apply the above two steps to the sub-array of elements with values smaller than the pivot and the sub-array of elements with values greater than the pivot.


TIME COMPLEXITY: O(nlog(n))
SPACE COMPLEXITY : O(1) using only stack space
*/

func QuickSort(arr *[]int, low, high int) {

	// if the array has more than one element than sorting will begin
	if low < high {
		// find the pivot
		pivot := pivotPosition(arr, low, high)

		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}

}

// find the pivot. and swap element's based on algorithm.
func pivotPosition(arr *[]int, low, high int) int {
	pivot := low
	i := low
	j := high

	for i < j {
		// For left.
		for (*arr)[i] <= (*arr)[pivot] && i <= high-1 {
			i++
		}

		// For right
		for (*arr)[j] >= (*arr)[pivot] && j >= low+1 {
			j--
		}

		// swap
		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[pivot], (*arr)[j] = (*arr)[j], (*arr)[pivot] // Pivot swap
	return pivot
}
