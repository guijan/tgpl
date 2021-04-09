// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    f, err := os.OpenFile("output.txt", os.O_CREATE | os.O_WRONLY, 0666)
    if (err != nil) {
        fmt.Fprintln(os.Stderr, "Error opening file")
        os.Exit(1)
    }
    for range os.Args[1:] {
        fmt.Fprintln(f, <-ch) // receive from channel ch
    }
    fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
    f.Close()
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
