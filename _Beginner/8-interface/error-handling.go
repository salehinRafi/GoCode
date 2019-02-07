package main

import "errors"
import "fmt"
import "math"

/*
	Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration.
*/

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		//Functions normally return error as last return value (string). Use errors.New to construct a basic error message.
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func ErrorHandling() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
