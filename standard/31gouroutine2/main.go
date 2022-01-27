package main

import (
	"fmt"
	"time"
)

func myGo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("name: %v\n", name)
		time.Sleep(time.Microsecond * 10)
	}
}
func main() {
	go myGo("1号协程")
	go myGo("2号协程")
	time.Sleep(time.Second)
}
