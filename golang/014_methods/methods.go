package main

import "fmt"

type Customer struct {
	id   int
	name string
}

func (customerobj Customer) String() string {
	return fmt.Sprintf("id: %d and name: %s", customerobj.id, customerobj.name)
}

func main() {
	customer1 := Customer{id: 10, name: "Sam"}
	fmt.Println(customer1.String())
}
