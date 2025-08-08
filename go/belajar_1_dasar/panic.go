package main

import "fmt"

func endApp()  {
	fmt.Println("Aplikasi Selesai")
	msg := recover()
	fmt.Println("Terjadi Panic, ",msg)
}

func runApp(err bool)  {
	defer endApp()
	if err {
		panic("Ups Error")
	}
}
func main()  {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma Lima = ", 3.5)

	runApp(true)
	fmt.Println("Lutfan Zainul Haq")
}