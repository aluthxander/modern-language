package main

import "fmt"

type User struct {
	City, Province, Country string
}

func main() {
	var alamat1 *User = new(User)
	var alamat2 *User = alamat1
	alamat2.City = "Bandung"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}