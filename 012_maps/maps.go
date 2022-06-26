package main

import "fmt"

func main() {
	// simpleMap()
	moreMap()
}

func simpleMap() {
	m := make(map[string]int)
	fmt.Println(m)
	m["one"] = 101
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Check point: ", m["one"])

	for v, k := range m { //The variable name has significance, v is always value, k is always key
		fmt.Println(k, v)
	}

	for v := range m {
		fmt.Println(v)
	}

	if v, ok := m["one"]; ok { //To check if a key is in the map. v -> value/zero value if key is not in map; ok is a boolean value, true-> if key is in the map
		fmt.Println("one: ", v)
	}

	if v, ok := m["two"]; ok { //To check if a key is in the map. v -> value/zero value if key is not in map; ok is a boolean value, true-> if key is in the map
		fmt.Println("two: ", v)
	}

}

func moreMap() {
	m2 := map[string]int{
		"a": 10,
		"b": 100,
		"c": 50,
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	delete(m2, "b")
	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
