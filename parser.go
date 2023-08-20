package main

import "errors"

type Parser struct {
	strSource []byte
	source    []Token
	curPos    int
	curToken  Token
}

func (psr *Parser) getNextToken() {
	psr.curPos++
	psr.curToken = psr.source[psr.curPos]
}

func (psr *Parser) peekToken() Token {
	if psr.curToken.code == tokens["EOF"] {
		return psr.curToken
	}
	return psr.source[psr.curPos+1]
}

func (psr *Parser) parse() (Token, error) {
	program := createToken(tokens["PROGRAM"], string(psr.strSource))

	psr.curToken = psr.source[psr.curPos]

	for psr.curToken.code != tokens["EOF"] {
		temp, err := psr.statement()
		if err != nil {
			return program, err
		}

		program.children = append(program.children, temp)
	}

	return program, nil
}

func (psr *Parser) statement() (Token, error) {
	tkn := createToken(tokens["STATEMENT"], "")
	switch psr.curToken.code {
	case tokens["FUN"]:
	case tokens["MUT"]:
	case tokens["CONST"]:
		tkn.children = append(tkn.children, createToken(tokens["CONST"], "const"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, errors.New("Parser Error: Expected \"TYPE\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, errors.New("Parser Error: Expected \"IDENT\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code == tokens["ASSIGN"] {
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			temp, err := psr.expression()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()
		}

		temp, err := psr.el()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

	case tokens["NOAS"]:
	case tokens["LOOP"]:
	case tokens["IF"]:
	default:
	}

	return tkn, nil
}

func (psr *Parser) expression() (Token, error) {
	tkn := createToken(tokens["EXPRESSION"], "")

	if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
		return tkn, errors.New("Parser Error: Expected \"IDENT\" or \"PRIMARY\"")
	}
	tkn.children = append(tkn.children, psr.curToken)

	for psr.peekToken().code == tokens["OPERATOR"] {
		psr.getNextToken()
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
			return tkn, errors.New("Parser Error: Expected \"IDENT\" or \"PRIMARY\"")
		}
	}

	return tkn, nil
}

func (psr *Parser) el() (Token, error) {
	if psr.curToken.code != tokens["EL"] {
		return psr.curToken, errors.New("Parser Error: Expected \"EL\"")
	}
	return psr.curToken, nil
}
