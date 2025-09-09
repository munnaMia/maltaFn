package sorting

func RecursiveInsertionSort(arr *[]int, i, length int) {
	if i == length {
		return
	}

	j := i
	for j > 0 && (*arr)[j-1] > (*arr)[j] {
		(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
		j--
	}

	RecursiveInsertionSort(arr, i+1, len((*arr)))

}
