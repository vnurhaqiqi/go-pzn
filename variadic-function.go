package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(8, 9 , 10)
	fmt.Println(total)

	slice := []int{2, 3, 4}
	totalTwo := sumAll(slice...)
	fmt.Println(totalTwo)
}