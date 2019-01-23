package main

import (
	"fmt"
)

func radix_sort() {

	slice := generatePositiveSlice(10)
	fmt.Println("\nRadix Sort:=")
	fmt.Println(slice)
	fmt.Println(RadixSort(slice))
}

// Finds the max number in an slice
func findMaxNum(slice []int) int {
	max := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

// Radix Sort
func RadixSort(slice []int) []int {
	// Base 10 is used
	max := findMaxNum(slice)
	size := len(slice)
	significantDigit := 1
	semiSorted := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for max/significantDigit > 0 {

		//fmt.Println("\tSorting: "+strconv.Itoa(significantDigit)+"'s place", slice)

		bucket := [10]int{0}

		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(slice[i]/significantDigit)%10]++
		}

		// Add the count of the previous buckets
		// Acquires the indexes after the end of each bucket location in the slice
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the bucket to fill a "semiSorted" slice
		for i := size - 1; i >= 0; i-- {
			bucket[(slice[i]/significantDigit)%10]--
			semiSorted[bucket[(slice[i]/significantDigit)%10]] = slice[i]
		}

		// Replace the current slice with the semisorted slice
		for i := 0; i < size; i++ {
			slice[i] = semiSorted[i]
		}

		//fmt.Println("\n\tBucket: ", bucket)

		// Move to next significant digit
		significantDigit *= 10
	}

	return slice
}
