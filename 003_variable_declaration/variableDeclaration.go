package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println(i)
	trial()
	variable_dec_def()
}

func trial() {
	y := 20
	fmt.Println(y)
}

func variable_dec_def() {
	var z int
	fmt.Println("Default value: ", z)
	z = 50
	fmt.Println(z)
}
