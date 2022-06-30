package main

import (
	"fmt"
	"mystring"
	"strings"
)

func main() {
	s := "This is a test 123 😯"
	s2 := strings.Map(myFunc, s)
	fmt.Println(s2)

	//Invoking a method from another package
	mystring.GetmyCard()
}

func myFunc(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}
	return in
}
