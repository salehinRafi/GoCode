package main

import (
	"fmt"
	"net/http"
)

func Pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	client := &http.Client{} //define method on pointer 
	resp, err := client.Get("https://www.google.com/")
	fmt.Println(resp, err)
}

/*
	Note that by default Go passes arguments by value (copying the arguments), if you want to pass the arguments by reference, you need to pass pointers.

	- To get the pointer of a value, use the & symbol in front of the value.
	- To dereference a pointer, use the * symbol.
	- Methods are often defined on pointers and not values (although they can be defined on both), so you will often store a pointer in a variable as in the `client` example above.
*/
