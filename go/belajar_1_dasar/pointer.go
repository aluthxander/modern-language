package main

import "fmt"

type User struct {
	City, Province, Country string
}
func main()  {
	addres1 := User{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(addres1)
	addres2 := &addres1
	addres2.City = "Bandung"
	fmt.Println(addres2)
}