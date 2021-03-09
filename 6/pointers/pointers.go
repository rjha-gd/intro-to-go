package main

import "fmt"

func zeroValue(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroValue(i)
	fmt.Println("zeroValue:", i)

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i)
	fmt.Println("Pointer:", &i)
}
