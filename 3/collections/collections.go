package main

import "fmt"

func main() {
	// arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slices
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:4])
	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// missing high index implies len(s)
	fmt.Println(mySlice[4:])

	// make slices
	cities := make([]string, 3)
	cities[0] = "Los Angeles"
	cities[1] = "San Diego"
	cities[2] = "San Jose"
	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))
	cities = append(cities, "San Francisco")
	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

	// range
	for i, city := range cities {
		fmt.Printf("%d: %s\n", i, city)
	}

	// map
	citiesWithPopulation := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	citiesWithPopulation["New York"] = 0
	for city, population := range citiesWithPopulation {
		fmt.Printf("%s has %d inhabitants\n", city, population)
	}

	delete(citiesWithPopulation, "Chicago")

	elem, ok := citiesWithPopulation["Chicago"]
	fmt.Printf("%v: %v\n", elem, ok)
}
