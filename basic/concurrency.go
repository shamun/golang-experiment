package main

import "fmt"

func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum
}

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c<- x:
            x, y = y, x + y
        case <- quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    // goroutine
    a := []int{7, 2, 10, -5, 9, 0, 4}
    c := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c

    fmt.Println(x, y, x + y)
    fmt.Println("------")

    // fibonacci
    fi := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-fi)
        }
        quit <- 0
    }()
    fibonacci(fi, quit)
}
