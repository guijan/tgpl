package main

import (
	"fmt"
	"strings"
    "os"
    "time"
)

func main() {
    echo1()
    echo2()
    echo3()
}


func echo1() {
	var s, sep string

    var startTime = time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

    fmt.Printf("echo1 took this long: %v\n", time.Since(startTime))
}

func echo2() {
    s, sep := "", ""
    var startTime = time.Now()
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
    fmt.Printf("echo2 took this long: %v\n", time.Since(startTime))
}

func echo3() {
    startTime := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Printf("echo3 took this long: %v\n", time.Since(startTime))
}
