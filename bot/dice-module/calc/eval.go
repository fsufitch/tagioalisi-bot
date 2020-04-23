package calc

import (
	"fmt"
	"strconv"

	srand "github.com/fsufitch/seedless-rand"
)

// DiceRoll encapsulates one roll of similar dice (e.g. 4d6)
type DiceRoll struct {
	Count   int
	Sides   int
	Results []int
}

// EvalResult is the numeric result of a postfix roll
type EvalResult struct {
	Value int
	Rolls []DiceRoll
}

func evalPostfix(tokens []Token) (EvalResult, error) {
	valStack := []int{}
	rolls := []DiceRoll{}

	for _, t := range tokens {
		switch {
		case t.Type == NumberToken:
			val, err := strconv.Atoi(t.String)
			if err != nil {
				return EvalResult{}, fmt.Errorf("invalid number: %s", t.String)
			}
			valStack = append(valStack, val)
		case t.Type != OperatorToken:
			return EvalResult{}, fmt.Errorf("invalid token: %+v", t)
		default:
			op, ok := ops[t.String]
			if !ok {
				return EvalResult{}, fmt.Errorf("unknown operator: %+v", t.String)
			}
			newValStack, newRolls, err := op(valStack)
			if err != nil {
				return EvalResult{}, err
			}
			valStack = newValStack
			rolls = append(rolls, newRolls...)
		}
	}

	if len(valStack) != 1 {
		return EvalResult{}, fmt.Errorf("not exactly 1 value in eval stack (%d)", len(valStack))
	}

	return EvalResult{Value: valStack[0], Rolls: rolls}, nil
}

var ops = map[string]func([]int) ([]int, []DiceRoll, error){
	"+": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for +")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		sum := x + y
		return append(rest, sum), nil, nil
	},
	"-": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for -")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		dif := x - y
		return append(rest, dif), nil, nil
	},
	"*": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for *")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		prod := x * y
		return append(rest, prod), nil, nil
	},
	"d": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for d")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		if x < 1 || x > 200 {
			return nil, nil, fmt.Errorf("invalid number of dice: %d", x)
		}
		if y < 2 {
			return nil, nil, fmt.Errorf("invalid number of sides for a die: %d", y)
		}
		sum, roll := rollDice(x, y)
		return append(rest, sum), []DiceRoll{roll}, nil
	},
	"min": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for min")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		min := x
		if y < x {
			min = y
		}
		return append(rest, min), nil, nil
	},
	"max": func(stack []int) ([]int, []DiceRoll, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("2 operands needed for max")
		}
		x := stack[len(stack)-2]
		y := stack[len(stack)-1]
		rest := stack[:len(stack)-2]
		max := x
		if y > x {
			max = y
		}
		return append(rest, max), nil, nil
	},
}

func rollDice(count, sides int) (int, DiceRoll) {
	sum := 0
	diceRoll := DiceRoll{Count: count, Sides: sides, Results: []int{}}
	for i := 0; i < count; i++ {
		roll := srand.Intn(sides) + 1
		sum += roll
		diceRoll.Results = append(diceRoll.Results, roll)
	}
	return sum, diceRoll
}
