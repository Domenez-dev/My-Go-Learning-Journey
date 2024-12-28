package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func calculate(a, b int, fn func(a, b int) int) int {
	fmt.Printf("Calculating... %v and %v \n", a, b)
	return fn(a, b)
}

func main() {
	result := calculate(1, 2, add)
	fmt.Println(result)

	result = calculate(1, 2, func(a, b int) int {
		return a - b
	})
	fmt.Println(result)
}
