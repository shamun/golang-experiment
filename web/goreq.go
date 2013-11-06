package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/franela/goreq"
)

type Item struct {
	Id   int
	Name string
}

func main() {

	item := Item{Id: 1111, Name: "foobar"}

	res, err := goreq.Request{
		Method: "POST",
		Uri:    "http://www.google.com",
		Body:   item,
	}.Do()

	if err != nil {
		fmt.Println("Something wrong ", err.Error())
	} else {
		spew.Dump(res.Header)
		fmt.Println(res.Body.ToString())
	}

}
