package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City: "Bandung",
		Province: "Jawa Barat",
		Country: "Indonesia",
	}

	address2 := &address1
	*address2 = Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
}
