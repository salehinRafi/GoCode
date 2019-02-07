package main

import "fmt"

func main() {
	/*
		Interfaces make the code more flexible, scalable and it’s a way to achieve polymorphism in Golang.
		Instead of requiring a particular type, interfaces allow to specify that only some behaviour is needed.

		1. Interface — set of methods required to implement such interface. It’s defined using keyword interface
		2. Interface type — variable of interface type which can hold any value implementing particular interface

	*/
	fmt.Println("------Interface Sections------")
	Interface()

	fmt.Println("------Interface Implicit Sections------")
	Implicit()

	/*
		Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration.
	*/
	fmt.Println("------Error Handling Interface Sections------")
	ErrorHandling()
}

/**

// define an interface  //
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_nameN [return_type]
}

// define a struct //
type struct_name struct {
   // .. variables .. //
}

// implement interface methods //
func (struct_name_variable struct_name) method_name1() [return_type] {
   // .. method implementation .. //
}

...

func (struct_name_variable struct_name) method_nameN() [return_type] {
   // .. method implementation .. //
}
*/
