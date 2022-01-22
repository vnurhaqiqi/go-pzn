package main

import "fmt"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodBye := sayGoodBye

	fmt.Println(goodBye("Sis"))
}