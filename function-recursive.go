package main

func recursiveFactorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * recursiveFactorial(num-1)
	}
}

func main() {
	recursive := recursiveFactorial(5)

	println(recursive)
}
