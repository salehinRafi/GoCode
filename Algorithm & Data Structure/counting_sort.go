/*
Counting Sort Program
	Counting Sort is O(n).
	It does not do any comparison. Instead, counting sort uses the actual values of the elements to index into an array.
	It only works for positive integers. The running time depends on the largest element. Therefore, if the maximum value is very large, the sorting takes long time.
*/

package main

import "fmt"

func counting_sort() {

	//slice := generatePositiveSlice(5)
	slice := []int{7, 1, 5, 2, 2}
	fmt.Println("\nCounting Sort:=")
	fmt.Println(slice)
	fmt.Println(CountingSort(slice))
}

// CountingSort assumes that each of the n input elements is an integer
// in the range 0 to k, for some integer k.
// 1. Create an array(slice) of the size of the maximum value + 1.
// 2. Count each element.
// 3. Add up the elements.
// 4. Put them back to result.
func CountingSort(slice []int) []int {

	// 1. Create an array(slice) of the size of the maximum value + 1
	k := GetMaxIntArray(slice)
	count := make([]int, k+1)

	// 2. Count each element
	for i := 0; i < len(slice); i++ {
		count[slice[i]] = count[slice[i]] + 1
	}

	//fmt.Println(count)

	// 3. Add up the elements
	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	//fmt.Println(count)

	// 4. Put them back to result
	result := make([]int, len(slice)+1)
	for j := 0; j < len(slice); j++ {
		result[count[slice[j]]] = slice[j]
		count[slice[j]] = count[slice[j]] - 1
	}

	return result
}

// Return the maximum value in an integer array(slice).
func GetMaxIntArray(slice []int) int {
	max_num := slice[0]
	for _, elem := range slice {
		//fmt.Printf("max_num: %v - elem: %v \n", max_num, elem)
		if max_num < elem {
			max_num = elem
		}
	}
	//fmt.Printf("final max_num: %v\n", max_num)
	return max_num
}
