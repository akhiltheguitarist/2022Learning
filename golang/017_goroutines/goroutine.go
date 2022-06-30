package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)   //number of go routines to run
	go func() { //closure
		runMe()
		wg.Done()
	}()
	go func() { //closure
		isRunning()
		wg.Done()
	}()
	wg.Wait()
}

func runMe() {
	fmt.Println("Started running")
}

func isRunning() {
	fmt.Println("Server is running")
}
