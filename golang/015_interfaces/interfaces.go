package main

import "fmt"

type customer interface {
	getFullName(int) string
}

func getFullName(a int) string {
	return "hello"
}

func main() {
	v := getFullName(50)
	fmt.Println(v)
}
