package main

import (
	"fmt"
	"sync"
)

func f(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go f("goroutine", &wg)

	wg.Add(1)
	go func(msg string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(msg)
	}("going", &wg)

	wg.Add(1)
	go f("direct", &wg)
	wg.Wait()

	fmt.Println("done")
}
