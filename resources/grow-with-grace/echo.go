package main

import (
    "io"
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
        //io.Copy(c, c)
        go io.Copy(c, c)  //concurrency
    }
}
