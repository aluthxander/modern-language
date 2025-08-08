package main

import "fmt"

func main() {
	var counter int = 0
	increment := func() int { counter++ ; return counter } // closure
	
	increment()
	increment()
	increment()
	fmt.Println(increment())
}