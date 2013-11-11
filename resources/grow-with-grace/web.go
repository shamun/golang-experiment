package main

import (
    "fmt"
    "log"
    "net/http"
)

const addr = "localhost:55556"

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal(err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, web!")
}
