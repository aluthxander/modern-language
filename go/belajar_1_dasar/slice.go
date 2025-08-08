package main

import "fmt"

func main() {
	names := [...]string{"Lutfan", "Zainul", "Haq", "Riyan", "Rozi"}
	slice := names[1:4]
	fmt.Println(slice)
	slice2 := names[1:]
	fmt.Println(slice2)
	slice3 := names[:3]
	fmt.Println(slice3)

	var slice4 []string = names[1:4]
	fmt.Println(slice4)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	slice5 := days[1:4]
	fmt.Println(slice5)
	slice5[0] = "Senin Lagi"
	fmt.Println(days)

	daySlice2 := append(slice5, "Jumat Lagi")
	daySlice2[0] = "Senin Lagi Lagi"
	fmt.Println(daySlice2)
	fmt.Println(days)
	// make slice
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "slice 0"
	newSlice[1] = "slice 1"
	newSlice = append(newSlice, "slice 2")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}