package main

type Parser struct {
	strSource []byte
	source    []Token
	curPos    int
	curToken  Token
	markers   []int
}

func (psr *Parser) setMarker() {
	psr.markers = append(psr.markers, psr.curPos)
}

func (psr *Parser) gotoMarker() {
	psr.curPos = psr.markers[len(psr.markers)-1]
	psr.curToken = psr.source[psr.curPos]
}

func (psr *Parser) removeMarker() {
	psr.markers = psr.markers[:len(psr.markers)-1]
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
		subTkn := createToken(tokens["DEFINITION"], "")

		subTkn.children = append(subTkn.children, createToken(tokens["FUN"], "fun"))
		psr.getNextToken()

		if psr.curToken.code == tokens["LBRACE"] {
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["TYPE"] {
				return tkn, createError("Parser", "statement -> FUN", "Expected \"TYPE\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code == tokens["REF"] {
				subTkn.children = append(subTkn.children, createToken(tokens["LBRACE"], "("))
				psr.getNextToken()
			}

			if psr.curToken.code != tokens["IDENT"] {
				return tkn, createError("Parser", "statement -> FUN", "Expected \"IDENT\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["RBRACE"] {
				return tkn, createError("Parser", "statement -> FUN", "Expected \"RBRACE\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()
		}

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> FUN", "Expected \"IDENT\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["LBRACE"] {
			return tkn, createError("Parser", "statement -> FUN", "Expected \"LBRACE\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		for psr.curToken.code == tokens["TYPE"] {
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["IDENT"] {
				return tkn, createError("Parser", "statement -> FUN", "Expected \"IDENT\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code == tokens["DELIMETER"] {
				if psr.peekToken().code != tokens["TYPE"] {
					return tkn, createError("Parser", "statement -> FUN", "Expected \"TYPE\"")
				}
				subTkn.children = append(subTkn.children, psr.curToken)
				psr.getNextToken()
			}
		}

		if psr.curToken.code != tokens["RBRACE"] {
			return tkn, createError("Parser", "statement -> FUN", "Expected \"RBRACE\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.block()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		tkn.children = append(tkn.children, subTkn)
	case tokens["MUT"]:
		subTkn := createToken(tokens["ASSIGNMENT"], "")

		subTkn.children = append(subTkn.children, createToken(tokens["MUT"], "mut"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> MUT", "Expected \"TYPE\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> MUT", "Expected \"IDENT\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code == tokens["ASSIGN"] {
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			temp, err := psr.expression()
			if err != nil {
				return tkn, err
			}
			subTkn.children = append(subTkn.children, temp)
			psr.getNextToken()
		}

		temp, err := psr.el()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		tkn.children = append(tkn.children, subTkn)
	case tokens["CONST"]:
		subTkn := createToken(tokens["ASSIGNMENT"], "")

		subTkn.children = append(subTkn.children, createToken(tokens["CONST"], "const"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"TYPE\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"IDENT\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> CONST", "Expected \"ASSIGN\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.expression()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		psr.getNextToken()

		temp, err = psr.el()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		tkn.children = append(tkn.children, subTkn)
	case tokens["NOAS"]:
		subTkn := createToken(tokens["ASSIGNMENT"], "")

		subTkn.children = append(subTkn.children, createToken(tokens["NOAS"], "noas"))
		psr.getNextToken()

		if psr.curToken.code != tokens["TYPE"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"TYPE\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["IDENT"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"IDENT\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> NOAS", "Expected \"ASSIGN\"")
		}
		subTkn.children = append(subTkn.children, psr.curToken)
		psr.getNextToken()

		temp, err := psr.expression()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		psr.getNextToken()

		temp, err = psr.el()
		if err != nil {
			return tkn, err
		}
		subTkn.children = append(subTkn.children, temp)
		tkn.children = append(tkn.children, subTkn)
	case tokens["IDENT"]:
		temp, err := psr.call()
		if err == nil {
			tkn.children = append(tkn.children, temp)
			psr.getNextToken()

			temp, err = psr.el()
			if err != nil {
				return tkn, err
			}
			tkn.children = append(tkn.children, temp)

			return tkn, nil
		}

		tkn.children = append(tkn.children, createToken(tokens["IDENT"], psr.curToken.text))
		psr.getNextToken()

		if psr.curToken.code != tokens["ASSIGN"] {
			return tkn, createError("Parser", "statement -> IDENT", "Expected \"ASSIGN\"")
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
	case tokens["LOOP"]:
		tkn.children = append(tkn.children, createToken(tokens["LOOP"], "loop"))
		psr.getNextToken()

		subTkn := createToken(tokens["ASSIGNMENT"], "")

		temp, err := psr.el()
		if err == nil {
			subTkn.children = append(subTkn.children, temp)
		} else {
			if psr.curToken.code != tokens["TYPE"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"TYPE\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["IDENT"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"IDENT\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			if psr.curToken.code != tokens["ASSIGN"] {
				return tkn, createError("Parser", "statement -> LOOP -> Assignment", "Expected \"ASSIGN\"")
			}
			subTkn.children = append(subTkn.children, psr.curToken)
			psr.getNextToken()

			temp, err = psr.expression()
			if err != nil {
				return tkn, err
			}
			subTkn.children = append(subTkn.children, temp)
			psr.getNextToken()

			temp, err = psr.el()
			if err != nil {
				return tkn, err
			}
			subTkn.children = append(subTkn.children, temp)
			psr.getNextToken()
		}

		tkn.children = append(tkn.children, subTkn)

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

func (psr *Parser) call() (Token, error) {
	tkn := createToken(tokens["CALL"], "")

	psr.setMarker()

	if psr.curToken.code != tokens["IDENT"] {
		psr.gotoMarker()
		psr.removeMarker()
		return tkn, createError("Parser", "call", "Expected \"IDENT\"")
	}
	tkn.children = append(tkn.children, psr.curToken)
	psr.getNextToken()

	if psr.curToken.code != tokens["LBRACE"] {
		psr.gotoMarker()
		psr.removeMarker()
		return tkn, createError("Parser", "call", "Expected \"LBRACE\"")
	}
	tkn.children = append(tkn.children, psr.curToken)
	psr.getNextToken()

	temp, err := psr.value()
	for err == nil {
		tkn.children = append(tkn.children, temp)
		psr.getNextToken()

		if psr.curToken.code != tokens["DELIMETER"] {
			if psr.curToken.code == tokens["RBRACE"] {
				break
			}

			psr.gotoMarker()
			psr.removeMarker()
			return tkn, createError("Parser", "call", "Expected \"DELIMETER\"")
		}
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err = psr.value()
	}

	if psr.curToken.code != tokens["RBRACE"] {
		psr.gotoMarker()
		psr.removeMarker()
		return tkn, createError("Parser", "call", "Expected \"RBRACE\"")
	}
	tkn.children = append(tkn.children, psr.curToken)

	psr.removeMarker()
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

	temp, err := psr.value()
	if err != nil {
		return tkn, createError("Parser", "expression", "Expected \"VALUE\"")
	}
	tkn.children = append(tkn.children, temp)

	for psr.peekToken().code == tokens["OPERATOR"] {
		psr.getNextToken()
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err = psr.value()
		if err != nil {
			return tkn, createError("Parser", "expression", "Expected \"VALUE\"")
		}
		tkn.children = append(tkn.children, temp)
	}

	return tkn, nil
}

func (psr *Parser) comparison() (Token, error) {
	tkn := createToken(tokens["COMPARISON"], "")

	temp, err := psr.value()
	if err != nil {
		return tkn, createError("Parser", "comparison", "Expected \"VALUE\"")
	}
	tkn.children = append(tkn.children, temp)

	for psr.peekToken().code == tokens["BOPERATOR"] {
		psr.getNextToken()
		tkn.children = append(tkn.children, psr.curToken)
		psr.getNextToken()

		temp, err = psr.value()
		if err != nil {
			return tkn, createError("Parser", "comparison", "Expected \"VALUE\"")
		}
		tkn.children = append(tkn.children, temp)
	}

	return tkn, nil
}

func (psr *Parser) el() (Token, error) {
	if psr.curToken.code != tokens["EL"] {
		return psr.curToken, createError("Parser", "el", "Expected \"EL\"")
	}
	return psr.curToken, nil
}

func (psr *Parser) value() (Token, error) {
	if psr.curToken.code == tokens["PRIMARY"] {
		return psr.curToken, nil
	} else if psr.curToken.code == tokens["IDENT"] {
		temp, err := psr.call()
		if err == nil {
			return temp, nil
		}
		return psr.curToken, nil
	} else {
		return psr.curToken, createError("Parser", "value", "Expected \"PRIMARY\" or \"IDENT\" or \"CALL\"")
	}
}
