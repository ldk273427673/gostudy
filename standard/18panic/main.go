package main

import "fmt"

func main() {
	// panic("crash")

	set_data(3)
	fmt.Println("everything is ok")
}

func set_data(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [2]int
	arr[x] = 20
}
