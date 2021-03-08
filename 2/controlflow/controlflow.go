package main

import "fmt"

func main() {
	// If
	a := 20
	if a%2 == 0 {
		fmt.Printf("%v is even\n", a)
	}

	// If with initializagtion
	if b := 10; b > 0 {
		fmt.Printf("%v is positive\n", b)
	}

	// If, Else If, If
	b := -1
	if b > 0 {
		fmt.Printf("%v is positive\n", b)
	} else if b == 0 {
		fmt.Printf("%v is zero\n", b)
	} else {
		fmt.Printf("%v is negative\n", b)
	}

	// For
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For with only condition
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// For without condition
	k := 0
	for {
		if k >= 5 {
			break
		}

		fmt.Println(k)
		k++
	}

	// switch
	x := 2
	fmt.Print("Write ", x, " as ")
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("?")
	}
}
