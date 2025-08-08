package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName)  {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string  {
	return person.Name
}

func main()  {
	person := Person{Name: "Lutfan Zainul Haq"}
	sayHello(person)
}