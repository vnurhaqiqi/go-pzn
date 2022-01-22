package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {
	filteredName := spamName(name)

	fmt.Println("Hello ", filteredName)
}

func spamName(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("viqi", spamName)
}