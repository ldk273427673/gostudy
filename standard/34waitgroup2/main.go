package main

import (
	"fmt"
	"sync"
)

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("worker %d:%d\n", x, i)
	}
}
func main() {
	var w sync.WaitGroup
	w.Add(2)
	go worker(1, &w)
	go worker(2, &w)
	w.Wait()
}
