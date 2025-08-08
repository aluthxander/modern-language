package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Lutfan Zainul Haq", "Zainul"))
	fmt.Println(strings.Split("Lutfan Zainul Haq", " "))
	fmt.Println(strings.ToLower("Lutfan Zainul Haq"))
	fmt.Println(strings.ToUpper("Lutfan Zainul Haq"))
	fmt.Println(strings.Trim("        Lutfan Zainul Haq       ", " "))
	fmt.Println(strings.ReplaceAll("Lutfan Zainul Haq", " ", "_"))
}