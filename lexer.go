package main

type Lexer struct {
	source  []byte
	curPos  int
	curChar byte
}

func (lxr Lexer) tokenize() []Token {
	lxr.curChar = lxr.source[lxr.curPos]

	return nil
}

func (lxr *Lexer) getNextToken() {

}

func (lxr *Lexer) getNextTokenNoWhiteSpace() {

}
