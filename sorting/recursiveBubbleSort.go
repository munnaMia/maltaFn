package sorting

func RecursiveBubbleSort(arr *[]int, length int) {
	if length == 1 {
		return
	}

	for j := 0; j < length-1; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
		}
	}

	RecursiveBubbleSort(arr, length-1)
}
