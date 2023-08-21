package main

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
		psr.getNextToken()
	}

	return program, nil
}

func (psr *Parser) statement() (Token, error) {
	tkn := createToken(tokens["STATEMENT"], "")
	switch psr.curToken.code {
	case tokens["FUN"]:
	case tokens["MUT"]:
		tkn.children = append(tkn.children, createToken(tokens["MUT"], "mut"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> MUT", "Expected \"TYPE\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> MUT", "Expected \"IDENT\"")
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
	case tokens["CONST"]:
		tkn.children = append(tkn.children, createToken(tokens["CONST"], "const"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"TYPE\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"IDENT\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"ASSIGN\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.expression()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

		temp, err = psr.el()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
	case tokens["NOAS"]:
		tkn.children = append(tkn.children, createToken(tokens["NOAS"], "noas"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"TYPE\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"IDENT\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"ASSIGN\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.expression()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

		temp, err = psr.el()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
	case tokens["IDENT"]:
		tkn.children = append(tkn.children, createToken(tokens["IDENT"], psr.curToken.text))
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> IDENT", "Expected \"ASSIGN\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.expression()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

		temp, err = psr.el()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
	case tokens["LOOP"]:
		tkn.children = append(tkn.children, createToken(tokens["LOOP"], "loop"))
		psr.getNextToken()

		temp, err := psr.el()
		if err == nil {
			tkn.children = append(tkn.children, temp)
		} else {
			if psr.curToken.code != tokens["TYPE"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"TYPE\"")
			}
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["IDENT"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"IDENT\"")
			}
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["ASSIGN"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"ASSIGN\"")
			}
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			temp, err = psr.expression()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()

			temp, err = psr.el()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()
		}

		temp, err = psr.el()
		if err == nil {
			tkn.children = append(tkn.children, temp)
		} else {
			temp, err = psr.comparison()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()

			temp, err = psr.el()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()
		}

		temp, err = psr.el()
		if err == nil {
			tkn.children = append(tkn.children, temp)
		} else {
			if psr.curToken.code != tokens["IDENT"] {
				return tkn, createError("Parser", "statement -> LOOP -> Modifier", "Expected \"IDENT\"")
			}
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["ASSIGN"] {
				return tkn, createError("Parser", "statement -> LOOP -> Modifier", "Expected \"ASSIGN\"")
			}
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			temp, err = psr.expression()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()

			temp, err = psr.el()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()
		}

		temp, err = psr.block()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
	case tokens["IF"]:
		tkn.children = append(tkn.children, createToken(tokens["IF"], psr.curToken.text))
		psr.getNextToken()

		temp, err := psr.comparison()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

		temp, err = psr.block()
		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)

		for psr.peekToken().code == tokens["ELIF"] {
			psr.getNextToken()
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			temp, err := psr.comparison()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()

			temp, err = psr.block()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
		}

		if psr.peekToken().code == tokens["ELSE"] {
			psr.getNextToken()
			tkn.children = append(tkn.children, psr.curToken)
			psr.getNextToken()

			temp, err = psr.block()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)
		}
	default:
		return tkn, createError("Parser", "statement", "Not a valid start to a statement")
	}

	return tkn, nil
}

func (psr *Parser) block() (Token, error) {
	tkn := createToken(tokens["BLOCK"], "")

	if psr.curToken.code != tokens["LSQUIRLY"] {
		return tkn, createError("Parser", "block", "Expected \"LSQUIRLY\"")
	}
	tkn.children = append(tkn.children, psr.curToken)
	psr.getNextToken()

	for psr.curToken.code != tokens["RSQUIRLY"] {
		if psr.curToken.code == tokens["EOF"] {
			return tkn, createError("Parser", "block", "\"EOF\" came before end of \"BLOCK\"")
		}

		temp, err := psr.statement()

		if err != nil {
			return tkn, err
		}
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()
	}

	tkn.children = append(tkn.children, psr.curToken)

	return tkn, nil
}

func (psr *Parser) expression() (Token, error) {
	tkn := createToken(tokens["EXPRESSION"], "")

	if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
		return tkn, createError("Parser", "expression", "Expected \"IDENT\" or \"PRIMARY\"")
	}
	tkn.children = append(tkn.children, psr.curToken)

	for psr.peekToken().code == tokens["OPERATOR"] {
		psr.getNextToken()
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
			return tkn, createError("Parser", "expression", "Expected \"IDENT\" or \"PRIMARY\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
	}

	return tkn, nil
}

func (psr *Parser) comparison() (Token, error) {
	tkn := createToken(tokens["COMPARISON"], "")

	if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
		return tkn, createError("Parser", "comparison", "Expected \"IDENT\" or \"PRIMARY\"")
	}
	tkn.children = append(tkn.children, psr.curToken)

	for psr.peekToken().code == tokens["BOPERATOR"] {
		psr.getNextToken()
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] && psr.curToken.code != tokens["PRIMARY"] {
			return tkn, createError("Parser", "comparison", "Expected \"IDENT\" or \"PRIMARY\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
	}

	return tkn, nil
}

func (psr *Parser) el() (Token, error) {
	if psr.curToken.code != tokens["EL"] {
		return psr.curToken, createError("Parser", "el", "Expected \"EL\"")
	}
	return psr.curToken, nil
}

func (psr *Parser) ref() (Token, error) {
	if psr.curToken.code != tokens["REF"] {
		return psr.curToken, createError("Parser", "ref", "Expected \"REF\"")
	}
	return psr.curToken, nil
}
