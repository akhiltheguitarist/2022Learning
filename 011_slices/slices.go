package main

import "fmt"

func main() {
	sliceBasic()
	sliceBasic2()
}

func sliceBasic() {
	s := make([]string, 0) //[]
	fmt.Println(s, len(s))
	s = append(s, "Sam") //[Sam]
	fmt.Println(s, len(s), s[0])
	s[0] = "Simon" //[Simon]
	fmt.Println(s, len(s), s[0])

	s2 := make([]string, 2) //[  ]
	s2 = append(s2, "Pap")  //Adds the value as 3rd element
	fmt.Println(s2, len(s2))
}

func sliceBasic2() {
	s3 := []string{"a", "b", "c"}
	for k, v := range s3 {
		fmt.Println(k, v) //k-> offset, v-> value
	}
	s3 = append(s3, "Mobi")
	for k, v := range s3 {
		fmt.Println(k, v)
	}
}
