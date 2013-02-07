package main

import (
    "fmt"
    "strings"
    "encoding/json"
)

const blob = `[
{"Title": "test!!", "URL": "http://golang.org"},
{"Title": "Keke", "URL": "http://google.com"}
]`

type Item struct {
    //Title string
    URL string
}

func main() {
    var items []*Item
    json.NewDecoder(strings.NewReader(blob)).Decode(&items)
    for _, item := range items {
        //fmt.Printf("Title: %v\tURL: %v\n", item.Title, item.URL)
        fmt.Printf("URL: %v\n", item.URL)
    }
}
