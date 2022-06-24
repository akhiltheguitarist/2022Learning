package main

import (
	"fmt"
)

func main() {
	simpleCondition()
	variableInCondition()
	forStatements()
	iteratingString()
}

func simpleCondition() {
	a := 15
	if a > 20 {
		fmt.Println("High value!")
	} else {
		fmt.Println("Below threshold")
	}
}

func variableInCondition() {
	a := "hello "
	if b := "world"; a == "pop" {
		fmt.Printf(a)
	} else {
		fmt.Println(a + b)
	}

	//Trying to access the variable b here causes error
	// fmt.Println(b)
}

func forStatements() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func iteratingString() {
	s := "My world"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
