package main

import (
	"fmt"
	"os"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		fmt.Println("No message provided")
	} else {
		fmt.Println(message)
	}
}
