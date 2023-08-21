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

	lexedSource := lexer.tokenize()
	for _, tkn := range lexedSource {
		fmt.Println(tkn.text)
		fmt.Println(tkn.code)
	}

	parser := Parser{
		dat,
		lexedSource,
		0,
		Token{},
	}

	parsedSource, err := parser.parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	parsedSource.print("")
}
