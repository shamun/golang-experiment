package main

import "fmt"

func main() {
	i := 0
	defer func() { fmt.Println(i) }()
	i++
	return
}
