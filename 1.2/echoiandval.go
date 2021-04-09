package main

import (
	"os"
	"fmt"
)

func main() {
	for i := 1; i < len(os.Args[:]); i++ {
		fmt.Printf("index %v: %v\n", i, os.Args[i])
	}
}
