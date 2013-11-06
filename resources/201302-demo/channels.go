package main

import "fmt"

func sum(a []int, c chan int) {
    sum := 0
    for i := 0; i < len(a); i++ {
        sum += a[i]
    }
    c <- sum
}

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int) // make(chan int, 6)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c

    fmt.Println(x, y, x + y)
}
