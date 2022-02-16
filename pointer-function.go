package main

import "fmt"

type Student struct {
	Name   string
	Status bool
}

func ChangeStatus(student *Student) {
	student.Status = false
}

func main() {
	student := Student{
		Name:   "Viqi",
		Status: true,
	}
	fmt.Println(student)

	ChangeStatus(&student)

	fmt.Println(student)
}
