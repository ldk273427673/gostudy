package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Printf("i的值是%d\n", i)
		}
		done <- true
	}()
	<-done
}
