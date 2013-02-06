package main

import (
    "fmt"
    "math/rand"
)

type binFunc func(int, int) int

func main() {
    (func() {
        fmt.Println("test!!!")
    })()

    // function collections
    fns := []binFunc{
        func(x, y int) int { return x+y},
        func(x, y int) int { return x-y},
        func(x, y int) int { return x*y},
    }

    fn := fns[rand.Intn(len(fns))]
    x, y := 5, 12
    fmt.Println(fn(x, y))
}
