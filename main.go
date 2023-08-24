package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

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

	parser := Parser{
		dat,
		lexedSource,
		0,
		Token{},
		[]int{},
	}

	parsedSource, err := parser.parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	//parsedSource.print("")

	semanticizer := Semanticizer{
		parsedSource,
	}

	err = semanticizer.checkValid(semanticizer.source, inBuiltFuncs)
	if err != nil {
		fmt.Println(err)
		return
	}

	optimizer := Optimizer{
		parsedSource,
	}

	optimizedSource := optimizer.optimize(optimizer.source)

	optimizedSource.print("")

	ellapsed := time.Since(start)
	fmt.Println("Compilation Completed in: " + strconv.Itoa(int(ellapsed.Milliseconds())) + "ms")
}
