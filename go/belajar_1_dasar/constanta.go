package main

import "fmt"

func main() {
	const name string = "Lutfan Zainul Haq"
	const hoby string = "Coding"
	fmt.Println(name)

	const (
		addres string = "Jakarta"
		hobby  string = "Coding"
	)
	fmt.Println(addres)
	fmt.Println(hobby)
}