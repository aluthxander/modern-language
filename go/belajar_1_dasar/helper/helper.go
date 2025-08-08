package helper
// jika diawali huruf kecil maka tidak bisa diakses di luar package tapi masih dapat diakses di dalam package
var version = "1.0.0"
var Application = "Belajar Go-Lang"

func sayGoodbye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}