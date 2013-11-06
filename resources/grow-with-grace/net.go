package main

import (
    "fmt"
    "log"
    "net"
)

const addr = "localhost:4000"

func main() {
    l, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Fprintln(c, "Hello")
        c.Close()
    }
}
