package main

import "fmt"

func factorialLoog(value int) int  {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoog(value - 1)
		
	}
}

func main() {
	result := factorialLoog(10)
	fmt.Println(result)
}