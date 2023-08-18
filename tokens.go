package main

type Token struct {
	code     int
	text     string
	children []Token
}

var tokens map[string]int = map[string]int{
	// Other
	"EOF":        0,
	"EL":         1,
	"IDENT":      2,
	"REF":        3,
	"OPERATOR":   4,
	"COMPARATOR": 5,
	"EXPRESSION": 6,
	"TYPE":       7,
	"PRIMARY":    8,
	"ILLEGAL":    9,
	"DELIMETER":  10,
	"ACCESSOR":   11,

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
}
