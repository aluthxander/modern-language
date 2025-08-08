package main

import "fmt"

func main() {
	type NoKtp string

	var ktpEko NoKtp = "111111111";

	var contoh string = "22222222222";
	var contohKtp NoKtp = NoKtp(contoh);

	fmt.Println(ktpEko);
	fmt.Println(contohKtp);
}