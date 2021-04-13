package main

import (
    "time"
    "fmt"
)

func main() {
    startTime := time.Now()
    PopCountShift(1234567890)
    fmt.Printf("PopCountShift took this long: %v\n", time.Since(startTime))
}

func PopCountShift(x uint64) int {
    var ret int

    for i := 0; i < 64; i++ {
        ret += int((x >> i) & 1)
    }
    return ret
}
