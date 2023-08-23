package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Token struct {
	code      int
	text      string
	children  []Token
	modifiers []string
}

func (token Token) print(indent string) {
	fmt.Println(indent + getTokenKey(token.code))
	for _, child := range token.children {
		child.print(indent + "    ")
	}
}

func createToken(code int, text string) Token {
	return Token{
		code,
		text,
		[]Token{},
		[]string{},
	}
}

func isKeyword(word string) bool {
	return slices.Contains(keyWords, word)
}

func isType(word string) bool {
	return slices.Contains(types, word)
}

func isPrimary(word string) bool {
	return slices.Contains(primaryWords, word)
}

func getTokenKey(code int) string {
	for key, val := range tokens {
		if code == val {
			return key
		}
	}
	return ""
}

var types []string = []string{
	"int",
	"bool",
	"string",
}

var keyWords []string = []string{
	"fun",
	"mut",
	"const",
	"noas",
	"loop",
	"if",
	"elif",
	"else",
	"cls",
	"stt",
}

var primaryWords []string = []string{
	"null",
	"false",
	"true",
}

var tokens map[string]int = map[string]int{
	// Other
	"EOF":        0,
	"ILLEGAL":    1,
	"INDECISIVE": 2,
	"EL":         3,
	"IDENT":      4,
	"REF":        5,
	"OPERATOR":   6,
	"COMPARATOR": 7,
	"TYPE":       8,
	"PRIMARY":    9,
	"DELIMETER":  10,
	"ACCESSOR":   11,
	"NOT":        12,
	"ASSIGN":     13,
	"BOPERATOR":  14,

	// Keywords
	"FUN":   20,
	"MUT":   21,
	"CONST": 22,
	"NOAS":  23,
	"LOOP":  24,
	"IF":    25,
	"ELIF":  26,
	"ELSE":  27,
	"CLS":   28,
	"STT":   29,
	"NULL":  30,

	// Parens
	"LSQUIRLY": 40,
	"RSQUIRLY": 41,
	"LBRACE":   42,
	"RBRACE":   43,
	"LSQUARE":  44,
	"RSQUARE":  45,
	"LANGLE":   46,
	"RANGLE":   47,

	// Parsing
	"PROGRAM":    100,
	"BLOCK":      101,
	"EXPRESSION": 102,
	"COMPARISON": 103,
	"STATEMENT":  104,
	"CALL":       105,
	"ASSIGNMENT": 106,
	"DEFINITION": 107,
}
