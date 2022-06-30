package main

import "fmt"

func main() {
	pointerAssign()
	pointerDeclare()
	passByReference()
}

func pointerAssign() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 50
	fmt.Println(a, b, *b, c)

	*b = 100
	fmt.Println(a, b, *b, c)

	c = 90
	fmt.Println(a, b, *b, c)
}

func pointerDeclare() {
	var b *int //default value will be nil

	fmt.Println("\n\n", b)
	// fmt.Println(*b)
}

func passByReference() {
	a := 100
	fmt.Println(a)
	getByReference(&a)
	fmt.Println(a)
}

func getByReference(a *int) {
	*a = 500
}
