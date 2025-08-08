package main

import "fmt"

func main() {
	var name[3] string
	
	name[0] = "Lutfan"
	name[1] = "Zainul"
	name[2] = "Haq"
	fmt.Println("Hello World")

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{10, 20, 30}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

	var values2 = [...]int{
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(values2)
}