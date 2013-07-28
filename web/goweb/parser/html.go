package main

import (
    "html/template"
    "os"
)

type Person struct {
    UserName string
}

func main() {
    t := template.New("fieldname")
    t, _ = t.Parse("hello {{.UserName}}!")
    p := Person{UserName: "kekeke"}
    t.Execute(os.Stdout, p)
}
