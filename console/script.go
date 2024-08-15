package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run script.go <string> <size>")
		return
	}
	strValue := os.Args[1]
	size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid size:", os.Args[2])
		return
	}
	for i := 0; i < size; i++ {
		fmt.Println(strValue)
	}
}