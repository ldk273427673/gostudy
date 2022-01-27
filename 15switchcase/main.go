package main

import "fmt"

func main() {
	edu := "本科"
	switch edu {
	case "本科":
		fmt.Println("你是本科学历")
	case "大专":
		fmt.Println("你是专科学历")
	}
}
