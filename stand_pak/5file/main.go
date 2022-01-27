package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("src/stand_pak/5file/1.txt")
	if err != nil {
		fmt.Println(err)
	}

	r := bufio.NewReader(file)

	buffer := make([]byte, 1024)
	for {
		i, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if i == 0 {
			break
		}
		fmt.Println(string(buffer[:i]))
	}
}
