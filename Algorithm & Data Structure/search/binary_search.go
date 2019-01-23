/*
Binary Search Program
	-	A binary search is a search strategy used to find elements within a list by consistently reducing the amount of data to be searched and thereby increasing the rate at which the search term is found.
	-	To use a binary search algorithm, the list to be operated on must have already been sorted.
*/
package main

import (
	"fmt"
)

func BinarySearch(target []int, value int) bool {

	start := 0
	end := len(target) - 1

	for start <= end {
		median := (start + end) / 2
		if target[median] < value {
			start = median + 1
		} else {
			end = median - 1
		}
	}
	if start == len(target) || target[start] != value {
		return false
	}
	return true
}

func binary_search() {

	values := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(BinarySearch(values, 4))

}
