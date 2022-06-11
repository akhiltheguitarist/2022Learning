package main

import "fmt"

func main() {
	s := "\"Hello\" \nthere \\!"
	fmt.Println(s)

	raw_string_invoke()
	string_combine()
	access_characters()
}

func raw_string_invoke() {
	raw_string := `"hello" world`
	fmt.Println("\n", raw_string)
}

func string_combine() {
	string1 := `Craig`
	string2 := `Simon`
	string3 := string1 + " & " + string2
	fmt.Println("\n", string3)
}

func access_characters() {
	s := "Hello world"
	b := s[0]
	c := s[5]
	fmt.Println("\n", s, b, string(b), c, string(c))
}
