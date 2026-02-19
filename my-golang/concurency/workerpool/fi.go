package main

import (
    "fmt"
    "sync"
)

// merge function combines multiple channels into a single channel
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.
    // Output copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are done.
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func main() {
    chan1 := make(chan int, 2)
    chan2 := make(chan int, 2)
    chan3 := make(chan int, 2)

    for i := 1; i <= 2; i++ {
        chan1 <- i
        chan2 <- i * 10
        chan3 <- i * 100
    }
    close(chan1)
    close(chan2)
    close(chan3)

    for n := range merge(chan1, chan2, chan3) {
        fmt.Println(n)
    }
}