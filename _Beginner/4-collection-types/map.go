package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// we mapping string type key to struct type value
var m map[string]Vertex

var n = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// Maps : Somewhat similar to what other languages call “dictionaries” or “hashes”.
// A map maps keys to values
// map[key-type]value-type is a key-value pair where we mapping type of key to type of value.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func Maps() {
	m = make(map[string]Vertex) // make function returns a map of the given type, initialized and ready for use.

	// Map Declaration :
	// 1. Map literals are like struct literals, but the KEYS are required.
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// 2. if the top-level type is just a type name, you can omit it from the elements of the literal.
	fmt.Println(n)           // get key-value pair
	fmt.Println(n["Google"]) // get specific value of key

	// Mutating Map
	o := make(map[string]int)

	o["Answer"] = 42 // Insert or update an element (value) :- m[key] = elem, , where m = map
	fmt.Println("The value:", o["Answer"])

	o["Answer"] = 48
	fmt.Println("The value:", o["Answer"])

	delete(o, "Answer") // Delete an element :- delete(m, key), where m = map
	fmt.Println("The value:", o["Answer"])

	v, ok := o["Answer"] //check if exist :- elem, ok = m[key], where ok is boolean
	fmt.Println("The value:", v, "Present?", ok)
}
