package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	id   int
	name string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	customer1 := customer{}
	fmt.Println(customer1) //{0 }

	customer2 := customer{102, "Chris"}
	fmt.Println(customer2) //{102 Chris}

	customer3 := customer{
		name: "Sam",
	}
	fmt.Println(customer3)
	customer3.id = 101
	fmt.Println(customer3)

	//Struct tags
	//Martialling data
	bob := `{"name":"Sam", "age":15}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
}
