package main

import "fmt"

type printer struct {
	message string
}

func (p *printer) SetMessage(msg string) {
	p.message = msg
}

func (p *printer) Print() string {
	return p.message
}

func main() {
	p := printer{}
	p.SetMessage("Start!")
	defer fmt.Println(p.Print())
	p.SetMessage("Done!")
}
