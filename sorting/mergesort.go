package sorting

/*
Merge sort is an efficient, divide-and-conquer sorting algorithm that works by recursively breaking down a list into sublists until each sublist consists of a single element, then merging those sublists in a sorted manner.

The core idea of merge sort is based on the principle that it's much easier to combine two sorted lists into a single sorted list than it is to sort a single unsorted list. It follows a three-step process:

Divide: The unsorted list is divided into n sublists, each containing one element. A list of one element is considered sorted.

Conquer (Recursion): The algorithm repeatedly merges sublists to produce new sorted sublists until there is only one sublist remaining.

Combine (Merge): The merging step is where the real work happens. Two sorted sublists are combined into a single sorted list by comparing the elements from each sublist one by one and placing the smaller element into the new list. This process is repeated until all elements from both sublists are moved to the new sorted list.

The time complexity of merge sort is O(n log n) in all cases (best, average, and worst), making it a very reliable choice for sorting. This consistent performance is due to its divide-and-conquer approach, which splits the problem into equal parts at each step.
*/
func MergeSort(arr *[]int, low, high int) {

	if low >= high {
		return
	}

	mid := int((high + low) / 2)

	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)

	merge(arr, low, mid, high)
}

func merge(arr *[]int, low, mid, high int) {

	//temp arr to store sorted elements
	temp := []int{}

	// two pointers
	left := low
	right := mid + 1

	for left <= mid && right <= high {

		// take smaller elements
		if (*arr)[left] <= (*arr)[right] {
			temp = append(temp, (*arr)[left])
			left++
		} else {
			temp = append(temp, (*arr)[right])
			right++
		}

	}

	// for leftover elements
	for left <= mid {
		temp = append(temp, (*arr)[left])
		left++
	}

	for right <= high {
		temp = append(temp, (*arr)[right])
		right++
	}

	// the trick
	// let's say low=3 & high = 7
	// where temp (0 to 4)
	// arr[i(low to high)] = temp[i - low]
	// arr[3]  = temp[3-3]
	// arr[4]  = temp[4-3]
	// arr[5]  = temp[5-3]
	// arr[6]  = temp[6-3]
	// arr[7]  = temp[7-3]
	// put the elements on main array
	for i := low; i <= high; i++ {
		(*arr)[i] = temp[i-low]
	}

}
