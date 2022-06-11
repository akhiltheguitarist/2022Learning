package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 8.785
	fmt.Println(i, f)

	type_conversion(i, f)
}

func type_conversion(x int8, y float32) {
	fmt.Println("Float: ", float32(x)+y)
	fmt.Println("Int: ", x+int8(y))
}
