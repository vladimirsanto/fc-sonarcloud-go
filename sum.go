package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
	fmt.Println(sub(10, 10))
	fmt.Println(mult(10, 10))
	fmt.Println(sumMult(10, 10))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b * a
}

func times(a int, b int) int {
	return a * b
}

func sumMult(a int, b int) int {
	return a + b*a
}
