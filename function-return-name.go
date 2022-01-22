package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Viqi"
	lastName = "Nurhaqiqi"

	return
}

func main() {
	a, b := getFullName()

	fmt.Println(a)
	fmt.Println(b)
}