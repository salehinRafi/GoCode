package main

import (
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(777)
	}
	return slice
}

// Generates a slice of size, size filled with positive numbers
func generatePositiveSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}
