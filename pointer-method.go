package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) AddMister() {
	man.Name = "Mr. " + man.Name
}

func main()  {
	viqi := Man{
		Name: "Viqi",
	}
	viqi.AddMister()

	fmt.Println(viqi)

}