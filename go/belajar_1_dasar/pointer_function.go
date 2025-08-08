package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(addres *Addres) {
	addres.Country = "Indonesia"
}

func main() {
	addresss := Addres{"Jakarta", "DKI Jakarta", "Mongolia"}
	fmt.Println(addresss)
	ChangeCountryToIndonesia(&addresss)
	fmt.Println(addresss)
}