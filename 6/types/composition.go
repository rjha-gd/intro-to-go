package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Id int
}

func main() {
	employee := Employee{
		Person: Person{Name: "Alice", Age: 42},
		Id:     1,
	}
	fmt.Println(employee)

	newEmployee := Employee{}
	newEmployee.Id = 2
	newEmployee.Name = "Bob"
	newEmployee.Age = 73
	fmt.Println(newEmployee)
	fmt.Println(newEmployee.Person)
}
