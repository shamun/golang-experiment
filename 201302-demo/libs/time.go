package main

import (
    "fmt"
    "time"
)

func main() {
    if time.Now().Hour() < 12 {
        fmt.Println("Good morning")
    } else {
        fmt.Println("Good afternoon")
    }

    home, err := time.Parse("2006-01-02 15:04", "2013-02-08 18:00")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(home)
    loc, _ := time.LoadLocation("Asia/Taipei")
    fmt.Println(home)
    fmt.Println(loc)
    fmt.Println(home.In(loc))
    age := time.Since(home.In(loc))
    fmt.Printf("Vacation time!! after %d hours", (age/time.Hour))
}
