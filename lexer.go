package main

import (
	"strings"
	"unicode"
)

type Lexer struct {
	source  []byte
	curPos  int
	curChar byte
}

func (lxr *Lexer) tokenize() []Token {
	tkns := []Token{}

	lxr.curChar = lxr.source[lxr.curPos]

	lxr.source = append(lxr.source, 3)

	for lxr.curChar != 3 {
		token := createToken(tokens["ILLEGAL"], "")

		switch lxr.curChar {
		case '`':
			for lxr.curChar != '\n' && lxr.curChar != '\r' {
				lxr.getNextChar()
			}
			lxr.getNextCharNoWhiteSpace()
			continue
		case '+':
			token = createToken(tokens["OPERATOR"], "+")
		case '-':
			token = createToken(tokens["OPERATOR"], "-")
		case '*':
			if lxr.getPeek() == '*' {
				lxr.getNextChar()
				token = createToken(tokens["OPERATOR"], "**")
			} else {
				token = createToken(tokens["OPERATOR"], "*")
			}
		case '/':
			if lxr.getPeek() == '/' {
				lxr.getNextChar()
				token = createToken(tokens["OPERATOR"], "//")
			} else {
				token = createToken(tokens["OPERATOR"], "/")
			}
		case '%':
			token = createToken(tokens["OPERATOR"], "%")
		case '&':
			if lxr.getPeek() == '&' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], "&&")
			} else {
				token = createToken(tokens["REF"], "&")
			}
		case '@':
			token = createToken(tokens["REF"], "@")
		case '!':
			if lxr.getPeek() == '=' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], "!=")
			} else {
				token = createToken(tokens["NOT"], "!")
			}
		case '=':
			if lxr.getPeek() == '=' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], "==")
			} else {
				token = createToken(tokens["ASSIGN"], "=")
			}
		case '|':
			if lxr.getPeek() == '|' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], "||")
			}
		case '<':
			if lxr.getPeek() == '=' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], "<=")
			} else {
				token = createToken(tokens["INDECISIVE"], "<")
			}
		case '>':
			if lxr.getPeek() == '=' {
				lxr.getNextChar()
				token = createToken(tokens["BOPERATOR"], ">=")
			} else {
				token = createToken(tokens["INDECISIVE"], ">")
			}
		case '(':
			token = createToken(tokens["LBRACE"], "(")
		case ')':
			token = createToken(tokens["RBRACE"], ")")
		case '{':
			token = createToken(tokens["LSQUIRLY"], "{")
		case '}':
			token = createToken(tokens["RSQUIRLY"], "}")
		case '[':
			token = createToken(tokens["LSQUARE"], "[")
		case ']':
			token = createToken(tokens["RSQUARE"], "]")
		case ';':
			token = createToken(tokens["EL"], ";")
		case '.':
			token = createToken(tokens["ACCESSOR"], ".")
		case ',':
			token = createToken(tokens["DELIMETER"], ",")
		}

		if unicode.IsLetter(rune(lxr.curChar)) {
			startPos := lxr.curPos
			for unicode.IsNumber(rune(lxr.getPeek())) || unicode.IsLetter(rune(lxr.getPeek())) {
				lxr.getNextChar()
			}
			word := lxr.source[startPos : lxr.curPos+1]
			if isKeyword(string(word)) {
				token = createToken(tokens[strings.ToUpper(string(word))], string(word))
			} else if isType(string(word)) {
				token = createToken(tokens["TYPE"], string(word))
			} else if isBool(string(word)) {
				token = createToken(tokens["BOOL"], string(word))
			} else {
				token = createToken(tokens["IDENT"], string(word))
			}
		} else if unicode.IsDigit(rune(lxr.curChar)) {
			startPos := lxr.curPos
			for unicode.IsNumber(rune(lxr.getPeek())) || lxr.getPeek() == '.' {
				lxr.getNextChar()
			}
			number := lxr.source[startPos : lxr.curPos+1]
			token = createToken(tokens["PRIMARY"], string(number))
		}

		tkns = append(tkns, token)

		lxr.getNextCharNoWhiteSpace()
	}

	tkns = append(tkns, createToken(tokens["EOF"], "EOF"))

	return tkns
}

func (lxr *Lexer) getNextChar() {
	lxr.curPos++
	lxr.curChar = lxr.source[lxr.curPos]
}

func (lxr Lexer) getPeek() byte {
	if lxr.curChar == 3 {
		return 3
	}
	return lxr.source[lxr.curPos+1]
}

func (lxr *Lexer) getNextCharNoWhiteSpace() {
	lxr.getNextChar()
	for lxr.curChar == '\r' || lxr.curChar == '\n' || lxr.curChar == ' ' || lxr.curChar == '\t' {
		lxr.getNextChar()
	}
}
