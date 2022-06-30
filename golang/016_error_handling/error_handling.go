package main

import (
	"errors"
	"fmt"
)

func main() {
	v, _ := modVal(20, 10)
	fmt.Println(v)
	v1, err := modVal(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v1)
	}

}

func modVal(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ERROR: Divisor can not be zero")
	}
	return a / b, nil
}
