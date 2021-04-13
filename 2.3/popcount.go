package main

import (
    "time"
    "fmt"
)

// pc[i] is the population count of i.
var pc[256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func main() {
    var testNum uint64 = 1234567890
    loopStart := time.Now()
    PopCountLoop(testNum)
    fmt.Printf("PopCountLoop time: %v\n", time.Since(loopStart))
    inlineStart := time.Now()
    PopCountInline(testNum)
    fmt.Printf("PopCountInline time: %v\n", time.Since(inlineStart))
}



// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
    var i int
    var ret int

    for i = 0; i < 64; i++ {
        ret += int(pc[byte(x>>i*8)])
    }
    return ret
}

// PopCountInline returns the population count (number of set bits) of x.
func PopCountInline(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}


