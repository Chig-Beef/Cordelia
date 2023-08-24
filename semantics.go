package main

import (
	"golang.org/x/exp/slices"
)

type Semanticizer struct {
	source Token
}

type Variable struct {
	Name      string
	Modifiers []string
	Type      string
	Variables []Variable
}

func (vr Variable) containsVar(name string) bool {
	for _, ident := range vr.Variables {
		if ident.Name == name {
			return true
		}
	}
	return false
}

func (smr *Semanticizer) checkValid(tkn Token, data map[string]Variable) error {
	var err error = nil

	switch tkn.code {
	case tokens["EOF"]:
	case tokens["ILLEGAL"]:
		err = createError("Semantics", "checkValid", "Cannot have \"ILLEGAL\" token")
	// INDECISIVE
	case tokens["EL"]:
	case tokens["IDENT"]:
		_, exists := data[tkn.text]
		if !exists {
			err = createError("Semantics", "checkValid", "\"IDENT\" was referenced before assignment ("+tkn.text+")")
		}
	case tokens["REF"]:
	case tokens["OPERATOR"]:
	case tokens["COMPARATOR"]:
	case tokens["TYPE"]:
	case tokens["PRIMARY"]:
	case tokens["DELIMETER"]:
	case tokens["ACCESSOR"]:
	case tokens["NOT"]:
	case tokens["ASSIGN"]:
	case tokens["BOPERATOR"]:
	case tokens["FUN"]:
	case tokens["MUT"]:
	case tokens["CONST"]:
	case tokens["NOAS"]:
	case tokens["LOOP"]:
	case tokens["IF"]:
	case tokens["ELIF"]:
	case tokens["ELSE"]:
	case tokens["CLS"]:
	case tokens["STT"]:
	case tokens["NULL"]:
	case tokens["LSQUIRLY"]:
	case tokens["RSQUIRLY"]:
	case tokens["LBRACE"]:
	case tokens["RBRACE"]:
	case tokens["LSQUARE"]:
	case tokens["RSQUARE"]:
	case tokens["LANGLE"]:
	case tokens["RANGLE"]:
	case tokens["PROGRAM"]:
	case tokens["BLOCK"]:
	case tokens["EXPRESSION"]:
	case tokens["COMPARISON"]:
	case tokens["STATEMENT"]:
	case tokens["CALL"]:
		err = smr.checkValid(tkn.children[0], data)
		if err != nil {
			return err
		}
		temp := data[tkn.children[0].text]
		if !slices.Contains(temp.Modifiers, "callable") {
			return createError("Semantics", "checkValid", "Attempted to call an uncallable \"IDENT\"")
		}
	case tokens["ASSIGNMENT"]:
		newVar := Variable{}
		if tkn.children[0].code == tokens["MUT"] {
			newVar.Modifiers = append(newVar.Modifiers, "change", "assign")
			newVar.Type = tkn.children[1].text
			data[tkn.children[2].text] = newVar
		} else if tkn.children[0].code == tokens["NOAS"] {
			newVar.Modifiers = append(newVar.Modifiers, "change")
			newVar.Type = tkn.children[1].text
			data[tkn.children[2].text] = newVar
		} else if tkn.children[0].code == tokens["CONST"] {
			newVar.Type = tkn.children[1].text
			data[tkn.children[2].text] = newVar
		} else if tkn.children[0].code == tokens["TYPE"] {
			newVar.Modifiers = append(newVar.Modifiers, "change", "assign")
			newVar.Type = tkn.children[0].text
			data[tkn.children[1].text] = newVar
		} else if tkn.children[0].code == tokens["IDENT"] {
			variable, exists := data[tkn.children[0].text]
			if !exists {
				err = createError("Semantics", "checkValid", "\"IDENT\" was referenced before assignment")
			} else if !slices.Contains(variable.Modifiers, "assign") {
				err = createError("Semantics", "checkValid", "Attempt was made to reassign an unassignable \"IDENT\"")
			}
		}
	case tokens["DEFINITION"]:
		if tkn.children[0].code == tokens["FUN"] {
			newVar := Variable{"", []string{"callable"}, "function", []Variable{}}

			index := 1
			if tkn.children[index].code == tokens["IDENT"] {
				newVar.Name = tkn.children[index].text
				data[tkn.children[index].text] = newVar
			} else {
				for tkn.children[index].code != tokens["RBRACE"] {
					index++
				}
				index++
				data[tkn.children[index].text] = newVar
			}
			index++
			for tkn.children[index].code != tokens["RBRACE"] {
				index++
				data[tkn.children[index+1].text] = Variable{tkn.children[index+1].text, []string{}, tkn.children[index].text, []Variable{}}
				index += 2
			}
		} else if tkn.children[0].code == tokens["STT"] {
			newVar := Variable{tkn.children[1].text, []string{"instantiates"}, "struct", []Variable{}}

			index := 2
			for tkn.children[index].code != tokens["RSQUIRLY"] {
				index++
				temp := Variable{tkn.children[index+1].text, []string{}, tkn.children[index].text, []Variable{}}
				newVar.Variables = append(newVar.Variables, temp)
				data[tkn.children[index+1].text] = temp
				index += 2
			}
			data[tkn.children[1].text] = newVar
		}
	case tokens["ACCESS"]:
		temp, exists := data[tkn.children[0].text]
		if !exists {
			err = createError("Semantics", "checkValid", "\"IDENT\" was referenced before assignment ("+tkn.text+")")
		}
		if len(temp.Variables) == 0 {
			err = createError("Semantics", "checkValid", "access was attempted on struct with no properties or methods ("+tkn.text+")")
		}
		if !temp.containsVar(tkn.children[2].text) {
			err = createError("Semantics", "checkValid", "\""+tkn.children[0].text+"\" does not contain property or method \""+tkn.children[2].text+"\"")
		}
	default:
		return createError("Semantics", "checkValid", "Token not recognised")
	}

	if err != nil {
		return err
	}

	for i := 0; i < len(tkn.children); i++ {
		err = smr.checkValid(tkn.children[i], data)
		if err != nil {
			return err
		}
	}
	return err
}
