package main

import (
    "fmt"
    "time"
)

func main() {
    startTime := time.Now()
    PopCountAnd(1234567890)
    fmt.Printf("PopCountAnd took this long: %v\n", time.Since(startTime))
}
func PopCountAnd(x uint64) {
    var ret int

    for ; x != 0; x &= x-1 {
        ret++
    }
}
