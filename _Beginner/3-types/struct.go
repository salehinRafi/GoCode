package main

import "fmt"

// There are no classes, only struct. Struct can have methods.
// Struct supports composition but not inheritance
// You don’t need to define getters and setters on struct fields, they can be accessed automatically
// BUT only exported fields (capitalized) can be accessed from outside of a package.

// Defining
type person struct {
	name string // lower case field, thus name is not accessible from outside a package.
	age  int
}

type secretAgent struct {
	person
	licenseToKill bool
}

//A struct literal sets a newly allocated struct value by listing the values of its fields.
// You can list just a subset of fields by using the "Name:" syntax (the order of named fields is irrelevant when using this syntax). The special prefix & constructs a pointer to a newly allocated struct.

// “initialize” a variable containing a struct or a reference by create a struct literals:
var (
	p = person{name: "Salehin Rafi", age: 1} // has type person
	q = &person{"Salehin Rafi", 1}           // has type *person
	r = person{name: "Salehin Rafi"}         // age:0 is implicit
	s = person{}                             // name:0 and age:0

	sa1 = secretAgent{ //  has type of secretAgent with person field and licenseToKill
		person{
			"Salehin",
			26,
		},
		true,
	}
)

// Struct : A struct is a collection of fields.
func Struct() {

	fmt.Println(p, q, r, s)

	//Accessing fields using the dot notation
	fmt.Println(p.name)
	fmt.Println(q.age)

	//Doing `q.age` is the same as doing `(*x).Name`, when `q` is a pointer.
	fmt.Println(q.age)
	fmt.Println((*q).age)

	//Common way to “initialize” a variable containing a struct or a reference to one, is to create a struct literal.
	//Another option is to create a constructor
	// use new expression to allocate a zeroed value of the requested type and to return a pointer to it.
	x := new(person)
	fmt.Println(*x)
}
