package calc

import "fmt"

var operatorPrecedences = map[string]int{
	"+":   10,
	"-":   10,
	"*":   20,
	"min": 30,
	"max": 30,
	"d":   40,
}

func infixToPostfix(infixTokens []Token) ([]Token, error) {
	// https://www.tutorialspoint.com/Convert-Infix-to-Postfix-Expression
	postfixTokens := []Token{}
	opstack := []Token{}

	for _, t := range infixTokens {
		switch {
		case t.Type == NumberToken:
			postfixTokens = append(postfixTokens, t)
		case t.Type != OperatorToken:
			return nil, fmt.Errorf("invalid token in infix/postfix conversion: %+v", t)
		case t.String == "(":
			opstack = append(opstack, t)
		case t.String == ")":
			for len(opstack) > 0 && opstack[len(opstack)-1].String != "(" {
				postfixTokens = append(postfixTokens, opstack[len(opstack)-1])
				opstack = opstack[:len(opstack)-1]
			}
			if len(opstack) == 0 {
				return nil, fmt.Errorf("unmatched close-parenthesis")
			}
			opstack = opstack[:len(opstack)-1]
		default:
			p, ok := operatorPrecedences[t.String]
			if !ok {
				return nil, fmt.Errorf("unknown operator: %s", t.String)
			}
			for len(opstack) > 0 && p <= operatorPrecedences[opstack[len(opstack)-1].String] {
				postfixTokens = append(postfixTokens, opstack[len(opstack)-1])
				opstack = opstack[:len(opstack)-1]
			}
			opstack = append(opstack, t)
		}
	}

	for i := len(opstack) - 1; i >= 0; i-- {
		if opstack[i].String == "(" {
			return nil, fmt.Errorf("unmatched open-parenthesis")
		}
		postfixTokens = append(postfixTokens, opstack[i])
	}

	return postfixTokens, nil
}
