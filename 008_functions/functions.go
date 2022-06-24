package main

import "fmt"

func main() {
	fmt.Println(addNumbers(4, 5))
	sum, div := addMultiplyNumber(10, 5)
	fmt.Println(sum, div)

	//To avoid the second returned value
	sum, _ = addMultiplyNumber(10, 5)
	fmt.Println(sum)

	//To avoid first returned value
	_, div = addMultiplyNumber(10, 5)
	fmt.Println(div)
	addMultiplyNumber(10, 5)
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func addMultiplyNumber(a int, b int) (int, int) {
	return a + b, a * b
}
