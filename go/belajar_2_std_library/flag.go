package main

import (
	"fmt"
	"flag"
)

func main() {
	host := flag.String("host", "localhost", "database host")
	user := flag.String("user", "root", "database user")
	password := flag.String("password", "", "database password")
	port := flag.Int("port", 3306, "database port")

	flag.Parse()
	
	fmt.Println("host:", *host)
	fmt.Println("user:", *user)
	fmt.Println("password:", *password)
	fmt.Println("port:", *port)
}