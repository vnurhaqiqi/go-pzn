package main

import "fmt"

func EndApp() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println("Aplikasi selesai")
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Aplikasi error")
	}
}

func main() {
	RunApp(true)
}