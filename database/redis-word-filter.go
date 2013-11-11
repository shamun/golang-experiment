package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"strings"
)

const (
	ADDRESS = "127.0.0.1:6379"
)

var (
	c, err = redis.Dial("tcp", ADDRESS)
)

func submitData(input []string) {
	if err != nil {
		log.Fatal(err)
	}

	c.Send("MULTI")

	c.Send("DEL", "user-words")

	c.Send("SADD", redis.Args{}.Add("user-words").AddFlat(input)...)

	c.Send("SINTER", "user-words", "bad-words")

	reply, err := c.Do("EXEC")

	if err != nil {
		fmt.Println(err)
	}

	values, _ := redis.Values(reply, nil)

	curse_words, err := redis.Strings(values[2], nil)
	if err != nil {
		fmt.Println(err)
	}

	if (len(curse_words)) > 0 {
		for _, v := range curse_words {
			fmt.Println(">>Found: ", v)
		}
	} else {
		fmt.Println(">>Nothing found")
	}
}

func main() {
	for {
		fmt.Println(">> Please enter some text with swear words than press Enter or \"q\" to exit")

		bio := bufio.NewReader(os.Stdin)
		line, _, _ := bio.ReadLine()

		if string(line) == "q" {
			break
		}

		terms := strings.Split(string(line), " ")
		submitData(terms)
	}

	fmt.Println("Session Ended!")
}
