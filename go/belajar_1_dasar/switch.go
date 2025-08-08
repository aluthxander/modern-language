package main

import "fmt"

func main() {
	name := "Lutfan"

	switch name {
	case "Lutfan":
		fmt.Println("Hello Lutfan")
	case "Zainul":
		fmt.Println("Hello Zainul")
	case "Haq":
		fmt.Println("Hello Haq")
	default:
		fmt.Println("Hello")
	}

	// short statment
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch dengan kondisi
	nilai := 80

	switch {
	case nilai > 75:
		fmt.Println("Selamat Anda Lulus")
	case nilai < 75:
		fmt.Println("Maaf Anda Tidak Lulus")
	default:
		fmt.Println("Nilai Tidak Diketahui")
	}
}