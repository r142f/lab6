package tokenizer

import (
	"fmt"
	"lab6/token"
	"regexp"
)

type Tokenizer struct {
	Input    string
	curChar  string
	curPos   int
	curToken token.Token
}

func (tokenizer *Tokenizer) Init() {
	if len(tokenizer.Input) == 0 {
		tokenizer.curChar = "$"
	} else {
		tokenizer.curChar = string(tokenizer.Input[0])
	}

	tokenizer.curPos++
}

func (tokenizer *Tokenizer) nextChar() {
	if tokenizer.curPos == len(tokenizer.Input) {
		tokenizer.curChar = "$"
	} else if tokenizer.curPos > len(tokenizer.Input) {
		panic(fmt.Errorf("can't execute nextChar(): end of input is reached"))
	} else {
		tokenizer.curChar = string(tokenizer.Input[tokenizer.curPos])
	}

	tokenizer.curPos++
}

func (tokenizer *Tokenizer) NextToken() (token.Token, error) {
	var err error
	defer func() {
		switch p := recover(); p.(type) {
		case error:
			err = p.(error)
		default:
			if p != nil {
				panic(p)
			}
		}
	}()

	re := regexp.MustCompile(`\s`)
	for re.MatchString(tokenizer.curChar) {
		tokenizer.nextChar()
	}

	switch tokenizer.curChar {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		token := token.NumberToken{Value: ""}

		re := regexp.MustCompile(`\d`)
		for re.MatchString(tokenizer.curChar) {
			token.Value += tokenizer.curChar
			tokenizer.nextChar()
		}

		tokenizer.curToken = &token

	case "(":
		tokenizer.nextChar()
		tokenizer.curToken = token.LEFT
	case ")":
		tokenizer.nextChar()
		tokenizer.curToken = token.RIGHT

	case "+":
		tokenizer.nextChar()
		tokenizer.curToken = token.ADD
	case "-":
		tokenizer.nextChar()
		tokenizer.curToken = token.SUB
	case "*":
		tokenizer.nextChar()
		tokenizer.curToken = token.MUL
	case "/":
		tokenizer.nextChar()
		tokenizer.curToken = token.DIV

	case "$":
		tokenizer.curToken = token.END

	default:
		err = fmt.Errorf("can't execute nextToken(): got illegal character at %v", tokenizer.curPos)
	}

	return tokenizer.curToken, err
}
