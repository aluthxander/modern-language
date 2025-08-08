package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age int
}
// method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main()  {
	var lutfan Customer
	lutfan.Name = "Lutfan Zainul Haq"
	lutfan.Addres = "Jakarta"
	lutfan.Age = 20

	fmt.Println(lutfan.Name)
	fmt.Println(lutfan.Addres)
	fmt.Println(lutfan.Age)
	fmt.Println(lutfan)

	// struct literal
	lutfan2 := Customer{"Lutfan Zainul Haq", "Semarang", 21}
	fmt.Println(lutfan2)
	lutfan3 := Customer{Name: "Lutfan Zainul Haq", Addres: "Surabaya", Age: 22}
	fmt.Println(lutfan3)

	lutfan3.sayHello("Anton")
}