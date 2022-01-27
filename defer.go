package main

import "fmt"

func Logging() {
	fmt.Println("Selesai memanggil fungsi")
}

func TestDefer(num int) {
	defer Logging()
	result := 10 / 2

	fmt.Println("Hasil ====>", result)
}

func main() {
	TestDefer(2)
}
