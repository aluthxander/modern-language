package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil fucntion")
}

func runApplication(){
	defer logging()
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApplication()
}