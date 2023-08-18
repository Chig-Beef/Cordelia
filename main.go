package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
	}

	lexer := Lexer{
		dat,
		0,
		0,
	}

	lexer.tokenize()
}
