package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	person := &Person{Name: name, Age: age}

	person = &Person{name, age}

	person = &Person{Name: name}
	person.Age = age
	return person
}

func (p *Person) String() string {
	return fmt.Sprintf("%v: %v", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	name := "John Doe"
	age := 42

	person := new(Person) // person := Person{}
	fmt.Println(person)
	person.Age = age
	person.Name = name
	fmt.Println(person)

	p := NewPerson(name, age)

	fmt.Println(p)
	fmt.Println(p.Age)
	fmt.Println(p.Name)

	p.Birthday()
	fmt.Println(p.Age)
}
