/*
Heap Sort Program
	-	HeapSort is a comparison-based sorting algorithm.
	-	HeapSort divides its input into a sorted and an unsorted region, and it iteratively shrinks the unsorted region by extracting the largest element and moving that to the sorted region. The improvement consists of the use of a heap data structure to find the maximun.
*/
package main

import (
	"fmt"
)

type MaxHeap struct {
	slice    []int
	heapSize int
}

func heap_sort() {
	slice := generateRandomSlice(11)
	fmt.Println("\nHeap Sort:=")
	fmt.Println(slice)
	fmt.Println(HeapSort(slice))
}

// turning slice into a max heap.
func BuildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	// Swap the first element of the list with the final element. Decrease the considered range of the list by one.
	if l < h.heapSize && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.heapSize && h.slice[r] > h.slice[max] {
		max = r
	}
	//log.Printf("MaxHeapify(%v): l,r=%v,%v; max=%v\t%v\n", i, l, r, max, h.slice)
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func HeapSort(slice []int) []int {
	h := BuildMaxHeap(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
	}
	return h.slice
}
