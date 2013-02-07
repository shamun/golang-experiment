package main

import (
    "fmt"
    "math/rand"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(time.Duration(rand.Intn(100)))
        fmt.Println(s)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    go say(" world")
    say("hello")
}
