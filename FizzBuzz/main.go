package main

import "fmt"

func main() {
	simple()
	optimize()
	for out := range routine(100) {
		fmt.Println(out)
	}

}

func simple() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func optimize() {
	for i := 1; i <= 100; i++ {
		result := ""

		if i%3 == 0 {
			result += "Fizz"
		}

		if i%5 == 0 {
			result += "Buzz"
		}

		if result != "" {
			fmt.Println(result)
			continue
		}
		fmt.Println(i)

	}
}

func routine(amount int) <-chan string {

	out := make(chan string, amount)

	go func() {
		for i := 1; i <= amount; i++ {
			result := ""
			if i%3 == 0 {
				result += "Fizz"
			}
			if i%5 == 0 {
				result += "Buzz"
			}
			if result == "" {
				result = fmt.Sprintf("%v", i)
			}
			out <- result
		}
		close(out)
	}()

	return out
}
