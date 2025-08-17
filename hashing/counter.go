package hashing

import "fmt"

// For each query,find out how many times the number appears in the array.
func Counter(arr, query []int) map[int]int {
	hash := make(map[int]int)

	// initialize the key's from query array and assign the with 0
	for i := range query {
		hash[query[i]] = 0
	}

	// counting there repeations
	for i := range arr {
		hash[arr[i]]++
	}
	return hash
}

// For each query, we need to find out how many times the character appears in the string.
func CharCounter(str string, query []rune) map[rune]int {
	hash := make(map[rune]int)
	runesSlice := []rune(str)

	// initialize the key's from query array and assign the with 0
	for i := range query {
		hash[query[i]] = 0
	}

	// counting there repeations
	for i := range runesSlice {
		hash[runesSlice[i]]++
	}
	return hash
}

// Takes a slice and tells which number has highest and lowest repeations
// if multiple loweset or hight numbers with same weight exist it will return random one.
func CounterHighandLowRepeats(arr []int) (int, int) {
	hash := make(map[int]int)

	// counting there repeations
	for i := range arr {
		hash[arr[i]]++
	}
	var maxKey, maxRpt, minKey, minRpt int
	// minRpt default value is 0 so it will create a runtime error to prevent this we intialized a value 
	minRpt = len(arr)
	fmt.Println(hash)
	for key, repeat := range hash {
		if maxRpt < repeat {
			maxKey = key   // key hold the value
			maxRpt = repeat // maximum repeation
		}
		if minRpt > repeat {
			minKey = key
			minRpt = repeat

		}
	}
	return maxKey, minKey
}
