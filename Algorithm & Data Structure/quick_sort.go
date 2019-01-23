/*
Quick Sort Program
	-	QuickSort is a Divide and Conquer algorithm.
	-	It picks an element as pivot and partitions the given array around the picked pivot.
	-	There are many different versions of quickSort that pick pivot in different ways:-
		i. 	 Always pick first element as pivot.
		ii.  Always pick last element as pivot (implemented below)
		iii. Pick a random element as pivot.
		iv.	 Pick median as pivot.
*/

package main

import (
	"fmt"
	"math/rand"
)

func quick_sort() {
	slice := generateRandomSlice(10)
	fmt.Println("\nQuick Sort:=")
	fmt.Println(slice)
	fmt.Println(QuickSort(slice))
}

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	// Set a random element as a pivot
	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
