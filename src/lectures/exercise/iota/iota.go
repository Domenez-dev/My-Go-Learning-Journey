//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import (
	"fmt"
)

const (
	add = iota
	sub
	mul
	div
)

type Operation int

func (operation Operation) calculate(a, b int) int {
	switch operation {
	case add:
		return a + b

	case sub:
		return a - b

	case mul:
		return a * b

	case div:
		return a / b
	default:
		return (0)
	}
}

func main() {
	add := Operation(add)
	sub := Operation(sub)
	mul := Operation(mul)
	div := Operation(div)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
