package main

import "fmt"

func main() {
    p := []int{1, 2, 33, 444, 22}
    fmt.Println(p[1:3])
    fmt.Println(p[:3])
    fmt.Println(p[2:])

    fmt.Println(p, len(p), cap(p))
}
