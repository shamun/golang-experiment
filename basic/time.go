package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	words := []string{"play", "golang", "is", "very", "interesting"}
	for i := len(words); i > 0; i-- {
		fmt.Print(rot13(words[i-1]), " ")
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
	fmt.Println()
}

func rot13(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c = (c-'A'+13)%26 + 'A'
		}
		if 'a' <= c && c <= 'z' {
			c = (c-'a'+13)%26 + 'a'
		}
		b[i] = c
	}
	return string(b)
}
