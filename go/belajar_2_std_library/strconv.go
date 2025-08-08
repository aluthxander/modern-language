package main

import (
	"fmt"
	"strconv"
)

func main()  {
	result, err := strconv.ParseFloat("10.234", 64)
	if err != nil  {
		panic(err)
	}
	fmt.Println(result)
}