package main

import "fmt"

func Ups() any {
	return "Ups"
}

func main() {
	var result any = Ups()
	fmt.Println(result)
}