package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("第一种，exec执行后，只看结果是否正确")
	cmd := exec.Command("go", "env")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() is failed ,the failed result is %v\n", err)
	} else {
		fmt.Println("the result is success!")
	}

}
