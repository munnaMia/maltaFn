package sorting

/*
Insertion sort is a simple sorting algorithm that builds the final sorted array one item at a time. It is efficient for small data sets but less efficient for large lists compared to more advanced algorithms like quicksort or merge sort. The basic idea is to virtually divide the array into a sorted and an unsorted part. Elements are then picked from the unsorted part and inserted into their correct position within the sorted part. 

Steps of Insertion Sort:
Assume the first element is sorted: The algorithm begins by considering the first element of the array as a sorted sub-list.
Pick the next element: Take the first element from the unsorted part of the array. This element is often referred to as the "key." 
Compare and shift: Compare the "key" with elements in the sorted sub-list, moving from right to left. If an element in the sorted sub-list is greater than the "key," shift it one position to the right to create space.
Insert the key: Once a position is found where the "key" is greater than or equal to the element to its left (or the beginning of the array is reached), insert the "key" into that position.
Repeat: Repeat steps 2-4 for all remaining elements in the unsorted part of the array until all elements have been moved to the sorted part.
*/

func InsertionSort(nums []int) []int{
	sorted := nums

	for i:=0 ; i<len(sorted);i++ {
		j := i
		for j> 0  && sorted[j-1] > sorted[j]{
			sorted[j-1],sorted[j] =sorted[j],sorted[j-1]
			j--
		}
	}

	return  sorted
}