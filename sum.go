package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
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

func sumMult2(a int, b int) int {
	return (a + b) * a
}

func subMult(a int, b int) int {
	return (a - b) * a
}
