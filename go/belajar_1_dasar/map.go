package main

import "fmt"

func main()  {
	var person map[string]string = map[string]string {}

	person["name"] = "Lutfan Zainul Haq"
	person["address"] = "Jakarta"

	fmt.Println(person)

	person2 := map[string]string {
		"name": "Lutfan Zainul Haq",
		"address": "Jakarta",
	}

	fmt.Println(person2["name"])

	book := map[string]string {
		"title": "Belajar Go-Lang",
		"author": "Lutfan Zainul Haq",
		"publisher": "Ebook",
	}

	book["abstrack"] = "Belajar Go-Lang dari zero to hero"

	fmt.Println(book)

	delete(book, "abstrack")

	fmt.Println(book)
}