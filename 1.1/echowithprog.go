package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var progname = os.Args[0] + ": "
	var args = strings.Join(os.Args[1:], " ")

	fmt.Println(progname + args)
}
