package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	port := os.Args[1]
	message := os.Args[2]

	fmt.Println(args)
}
