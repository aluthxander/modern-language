package main

import "fmt"

func main() {
	var niali32 int32 = 3761552
	var niali64 int64 = int64(niali32)
	var nilai16 int16 = int16(niali32)

	fmt.Println(niali32)
	fmt.Println(niali64)
	fmt.Println(nilai16)

	var name = "Lutfan"
	var karakter = name[0]
	var eString string = string(karakter)
	fmt.Println(karakter)
	fmt.Println(eString)
}