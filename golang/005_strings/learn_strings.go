package main

import "fmt"

func main() {
	s := "\"Hello\" \nthere \\!"
	fmt.Println(s)

	raw_string_invoke()
	string_combine()
	access_characters()
	substringemoji()
	arraycheck()
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
	c := s[7]
	fmt.Println("\n", s, b, string(b), c, string(c))
	fmt.Printf("Length of string: %v", len(s))
}

func substringemoji() {
	s := "😊❤️"
	a := s[:1]
	b := s[2:]
	fmt.Println("\n\n", s, "\t", a, "\t", b)
}

func arraycheck() {
	var myArr [2]string
	myArr[0] = "10"
	fmt.Println(myArr)

	var myArrInt [2]int
	myArrInt[0] = 10
	fmt.Println(myArrInt)
}
