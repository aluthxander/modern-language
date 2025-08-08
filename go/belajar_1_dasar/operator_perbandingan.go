package main

import "fmt"

func main() {
	var name1 = "Lutfan Zainul Haq"
	var name2 = "Lutfan zainul Haq"

	var result bool = name1 == name2
	fmt.Println(result)
}