package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":55555", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[1:]
    fmt.Fprint(w, "Hello, " + name)
}
