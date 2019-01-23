/*
Merge Sort Program
	-	Merge Sort is a Divide and Conquer algorithm.
	-	It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.
*/

package main

import (
	"fmt"
)

func merge_sort() {
	slice := generateSlice(10)
	fmt.Println("\nMerge Sort:=")
	fmt.Println(slice)
	fmt.Println(MergeSort(slice))
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {
	ret := make([]int, 0, len(left)+len(right)) // Creating slice with specify capacity
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right...)
		}
		if len(right) == 0 {
			return append(ret, left...)
		}
		// keep track of the left and right slices to be merged and just pick the next smallest value from the left or right slice to add the new slice
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	return ret

}

// Run MergeSort algorithm on a slice single using top-down approach
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	// Divide slice into two
	mid := len(slice) / 2

	left := MergeSort(slice[:mid])  //recursively call left portions of slice
	right := MergeSort(slice[mid:]) //recursively call right portions of slice
	return Merge(left, right)
}
