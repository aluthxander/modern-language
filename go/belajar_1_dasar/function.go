package main

import "fmt"


func sayHello()  {
	fmt.Println("Hello World")
}

// function dengan parameter
func sayHelloWithName(name string)  {
	fmt.Println("Hello", name)
}

// function dengan return
func total(a int, b int) int {
	return a + b
}

// function dengan multiple return
func total2(a int, b int) (int, int) {
	return a + b, a - b
}

// variadic function
func sumAll(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
// type declaration
type Filter func(string) string


// function dengan parameter dan return
func getGoodByeFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Good Bye ", filteredName)
}
// function value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello()
	sayHelloWithName("Lutfan")

	hasil := total(5, 10)
	fmt.Println(hasil)

	fmt.Println(total2(5, 10))

	hasilTambah, hasilKurang := total2(5, 10)
	_, kurang := total2(20, 17)
	fmt.Println("Total Tambah:", hasilTambah, "Total Kurang:", hasilKurang)
	fmt.Println("Kurang:", kurang)

	sumAngka := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total Angka:", sumAngka)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumAngka2 := sumAll(numbers...)
	fmt.Println("Total Angka:", sumAngka2)

	contoh := getGoodBye
	fmt.Println(contoh("Lutfan"))

	getGoodByeFilter("Lutfan", spamFilter)
	getGoodByeFilter("Anjing", spamFilter)

	// anonymus function
	getGoodByeFilter("Lutfan", func(name string) string {
		return "Good Bye " + name
	})
}