package main

import (
	"errors"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) (z int) {
	z = x * y
	return
}

func divide(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func safeDivide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	q := x / y
	r := x % y
	return q, r, nil
}

func main() {
	sum := add(2, 3)
	fmt.Printf("Sum: %v\n", sum)
	difference := subtract(2, 3)
	fmt.Printf("Difference: %v\n", difference)
	product := multiply(2, 3)
	fmt.Printf("Product: %v\n", product)
	quotient, remainder := divide(14, 3)
	fmt.Printf("Quotient: %v, Remainder: %v\n", quotient, remainder)

	safeQuotient, safeRemainder, err := safeDivide(14, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Quotient: %v, Remainder: %v\n", safeQuotient, safeRemainder)
	}
}
