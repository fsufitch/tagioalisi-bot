package calc

import (
	"io"
	"strings"

	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/pkg/errors"
)

// DiceCalculator encapsulates all logic for calculating dice rolls
type DiceCalculator struct {
	Log *log.Logger
}

// CalculatorResult is a result from a query to DiceCalculator
type CalculatorResult struct {
	EvalResult
	InfixTokens   []Token
	PostfixTokens []Token
}

// Calculate is the entrypoint to dice rolling/calculation
func (c DiceCalculator) Calculate(input string) (CalculatorResult, error) {
	c.Log.Debugf("dice: calculating %s", input)
	reader := strings.NewReader(input)
	tokenizer := NewTokenizer(reader)

	infixTokens := []Token{}
	for {
		token, err := tokenizer.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return CalculatorResult{}, errors.Wrapf(err, "could not tokenize `%s`", input)
		}
		infixTokens = append(infixTokens, token)
	}

	c.Log.Debugf("dice: infix-tokenized to %+v", infixTokens)

	postfixTokens, err := infixToPostfix(infixTokens)
	if err != nil {
		return CalculatorResult{}, errors.Wrapf(err, "could not convert to postfix `%s`", infixTokens)
	}

	c.Log.Debugf("dice: converted to postfix %+v", postfixTokens)

	result, err := evalPostfix(postfixTokens)
	if err != nil {
		return CalculatorResult{}, errors.Wrapf(err, "could not evaluate postfix `%s`", postfixTokens)
	}

	c.Log.Debugf("dice: evaluated result %+v", result)
	return CalculatorResult{EvalResult: result, InfixTokens: infixTokens, PostfixTokens: postfixTokens}, nil
}
