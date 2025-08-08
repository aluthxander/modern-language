package main

import "fmt"

func main() {
	var nilai = 80
	var absensi = 60

	if nilai > 75 && absensi > 75 {
		fmt.Println("Selamat Anda Lulus")
	} else {
		fmt.Println("Maaf Anda Tidak Lulus")
	}

	if length := len("Haq"); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	}else{
		fmt.Println("Nama sudah benar")
	}
}