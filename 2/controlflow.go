package main

import "fmt"

func main() {
	a := 20
	if a%2 == 0 {
		fmt.Printf("%v is even\n", a)
	}

	if b := 10; b > 0 {
		fmt.Printf("%v is positive\n", b)
	}

	b := -1
	if b > 0 {
		fmt.Printf("%v is positive\n", b)
	} else if b == 0 {
		fmt.Printf("%v is zero\n", b)
	} else {
		fmt.Printf("%v is negative\n", b)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		if k >= 5 {
			break
		}

		fmt.Println(k)
		k++
	}
}
