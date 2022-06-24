package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	switch word {
	case "hello":
		fmt.Println(word)
	case "pop":
		fmt.Println(word + "peep")
	default:
		fmt.Println("Good day")
	}
}
