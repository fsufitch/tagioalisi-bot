package calc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// TokenType is an int to categorize tokenizer tokens
type TokenType int

// Enumeration of the types of tokens the tokenizer can produce
const (
	InvalidToken TokenType = iota
	NumberToken
	OperatorToken
)

// Token is a parsed out token
type Token struct {
	Type   TokenType
	String string
}

// Tokenizer is an object that can break up an input into tokens
type Tokenizer interface {
	Next() (Token, error)
}

type tokenizer struct {
	scanner *bufio.Scanner
}

// NewTokenizer creates a new Tokenizer from a string
func NewTokenizer(input io.Reader) Tokenizer {
	sc := bufio.NewScanner(input)
	sc.Split(ReadTokens)
	return tokenizer{
		scanner: sc,
	}
}

func (t tokenizer) Next() (Token, error) {
	if !t.scanner.Scan() {
		if err := t.scanner.Err(); err != nil {
			return Token{}, err
		}
		return Token{}, io.EOF
	}

	text := t.scanner.Text()
	switch text {
	case "(", ")", "+", "-", "*", "d", "min", "max":
		return Token{Type: OperatorToken, String: text}, nil
	}

	if _, err := strconv.Atoi(text); err == nil {
		return Token{Type: NumberToken, String: text}, nil
	}

	return Token{}, fmt.Errorf("unexpected token: %s", text)
}

// ReadTokens implements bufio.SplitFunc for tokenizing the input
func ReadTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	skip := 0
	runes := []rune{}
	for {
		r, size := utf8.DecodeRune(data[skip:])
		if r == utf8.RuneError {
			if size == 0 {
				// Out of data
				if atEOF {
					// Skip everything remaining, finish up
					return skip, nil, nil
				}
				// Request more data
				return 0, nil, nil
			}
			return skip, nil, errors.New("invalid utf-8")
		}

		skip += size
		if !unicode.IsSpace(r) {
			runes = append(runes, r)
			break
		}
	}

	switch runes[0] {
	case '(', ')', '+', '-', '*':
		return skip, []byte(string(runes)), nil
	}

	mode := 0 // 1=letters, 2=numbers
	if unicode.IsLetter(runes[0]) {
		mode = 1
	} else if unicode.IsNumber(runes[0]) {
		mode = 2
	} else {
		return 0, nil, fmt.Errorf("unexpected symbol, expected operator, letter or number: %s", string(runes[0]))
	}

	for {
		r, size := utf8.DecodeRune(data[skip:])
		if r == utf8.RuneError {
			if size == 0 {
				// Out of data
				if atEOF {
					// Return what we got, finish up
					return skip, []byte(string(runes)), nil
				}
				// Ask for more data
				return 0, nil, nil
			}
			return skip, nil, errors.New("invalid utf-8")
		}

		if (mode == 1 && !unicode.IsLetter(r)) || // Finished a word token
			(mode == 2 && !unicode.IsNumber(r)) { // Finished a number token
			return skip, []byte(string(runes)), nil
		}

		skip += size
		runes = append(runes, r)
	}
}
