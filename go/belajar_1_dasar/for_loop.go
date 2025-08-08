package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Selesai")
	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan 2 ke", i)
	}

	fmt.Println("Selesai")
	// for untuk data collection
	names := []string{"Lutfan", "Zainul", "Haq"}
	for i, name := range names {
		fmt.Println("Perulangan 3 ke", i, name)
	}
	fmt.Println("Selesai")
	for _, name := range names {
		fmt.Println("Perulangan 3 ke", name)
	}

	// continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("Selesai")

	// break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}