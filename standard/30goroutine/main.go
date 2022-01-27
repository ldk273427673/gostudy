package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("hello goroutine")
}

func main() {
	go test()
	fmt.Println("hello mainfunc")
	time.Sleep(time.Second * 2)
}
