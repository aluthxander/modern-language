package main

import "fmt"

func main() {
	var angka1 = 10
	var angka2 = 20
	var angka3 = angka1 + angka2 - 15
	fmt.Println(angka3)

	var i = 10
	i+= 12
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}