package main

import (
	"golang.org/x/exp/slices"
)

type Semanticizer struct {
	source Token
}

func (smr *Semanticizer) checkValid(tkn Token, data map[string][]string) error {
	var err error

	switch tkn.code {
	case tokens["EOF"]:
		err = nil
	case tokens["ILLEGAL"]:
		err = createError("Semantics", "checkValid", "Cannot have \"ILLEGAL\" token")
	// INDECISIVE
	case tokens["EL"]:
		err = nil
	case tokens["IDENT"]:
		err = nil
		_, exists := data[tkn.text]
		if !exists {
			err = createError("Semantics", "checkValid", "\"IDENT\" was referenced before assignment.")
		}
	case tokens["REF"]:
		err = nil
	case tokens["OPERATOR"]:
		err = nil
	case tokens["COMPARATOR"]:
		err = nil
	case tokens["TYPE"]:
		err = nil
	case tokens["PRIMARY"]:
		err = nil
	case tokens["DELIMETER"]:
		err = nil
	case tokens["ACCESSOR"]:
		err = nil
	case tokens["NOT"]:
		err = nil
	case tokens["ASSIGN"]:
		err = nil
	case tokens["BOPERATOR"]:
		err = nil
	case tokens["FUN"]:
		err = nil
	case tokens["MUT"]:
		err = nil
	case tokens["CONST"]:
		err = nil
	case tokens["NOAS"]:
		err = nil
	case tokens["LOOP"]:
		err = nil
	case tokens["IF"]:
		err = nil
	case tokens["ELIF"]:
		err = nil
	case tokens["ELSE"]:
		err = nil
	case tokens["CLS"]:
		err = nil
	case tokens["STT"]:
		err = nil
	case tokens["NULL"]:
		err = nil
	case tokens["LSQUIRLY"]:
		err = nil
	case tokens["RSQUIRLY"]:
		err = nil
	case tokens["LBRACE"]:
		err = nil
	case tokens["RBRACE"]:
		err = nil
	case tokens["LSQUARE"]:
		err = nil
	case tokens["RSQUARE"]:
		err = nil
	case tokens["LANGLE"]:
		err = nil
	case tokens["RANGLE"]:
		err = nil
	case tokens["PROGRAM"]:
		err = nil
	case tokens["BLOCK"]:
		err = nil
	case tokens["EXPRESSION"]:
		err = nil
	case tokens["COMPARISON"]:
		err = nil
	case tokens["STATEMENT"]:
		err = nil
	case tokens["CALL"]:
		err = smr.checkValid(tkn.children[0], data)
		if err != nil {
			return err
		}
		variable := data[tkn.children[0].text]
		if !slices.Contains(variable, "callable") {
			return createError("Semantics", "checkValid", "Attempted to call an uncallable \"IDENT\"")
		}
	case tokens["ASSIGNMENT"]:
		err = nil
		newVar := []string{}
		if tkn.children[0].code == tokens["MUT"] {
			newVar = append(newVar, "change", "assign")
		} else if tkn.children[0].code == tokens["NOAS"] {
			newVar = append(newVar, "change")
		}
		newVar = append(newVar, "type:"+tkn.children[1].text)
		data[tkn.children[2].text] = newVar
	case tokens["DEFINITION"]:

		err = nil
		newVar := []string{"callable"}

		index := 1
		if tkn.children[index].code == tokens["IDENT"] {
			data[tkn.children[index].text] = newVar
		} else {
			index++
			for tkn.children[index].code != tokens["RBRACE"] {
				index++
			}
			index++
			data[tkn.children[index].text] = newVar
		}
	default:
		return createError("Semantics", "checkValid", "Token not recognised")
	}

	for i := 0; i < len(tkn.children); i++ {
		err = smr.checkValid(tkn.children[i], data)
		if err != nil {
			return err
		}
	}
	return err
}
