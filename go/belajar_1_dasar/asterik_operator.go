package main

import "fmt"

type User struct {
	City, Province, Country string
}
func main()  {
	addres1 := User{"Jakarta", "DKI Jakarta", "Indonesia"}
	addres2 := &addres1
	addres2.City = "Bandung"
	fmt.Println(addres1)
	fmt.Println(addres2)

	addres2 = &User{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(addres2)
}