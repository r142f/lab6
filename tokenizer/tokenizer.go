package tokenizer

import (
	"fmt"
	"lab6/token"
	"regexp"
)

type Tokenizer struct {
	Input    string
	CurChar  string
	CurPos   int
	CurToken token.Token
}

func (tokenizer *Tokenizer) Init() {
	if len(tokenizer.Input) == 0 {
		tokenizer.CurChar = "$"
	} else {
		tokenizer.CurChar = string(tokenizer.Input[0])
	}

	tokenizer.CurPos++
}

func (tokenizer *Tokenizer) nextChar() {
	if tokenizer.CurPos == len(tokenizer.Input) {
		tokenizer.CurChar = "$"
	} else if tokenizer.CurPos > len(tokenizer.Input) {
		panic(fmt.Errorf("can't execute nextChar(): end of input is reached"))
	} else {
		tokenizer.CurChar = string(tokenizer.Input[tokenizer.CurPos])
	}

	tokenizer.CurPos++
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
	for re.MatchString(tokenizer.CurChar) {
		tokenizer.nextChar()
	}

	switch tokenizer.CurChar {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		token := token.NumberToken{Value: ""}

		re := regexp.MustCompile(`\d`)
		for re.MatchString(tokenizer.CurChar) {
			token.Value += tokenizer.CurChar
			tokenizer.nextChar()
		}

		tokenizer.CurToken = &token

	case "(":
		tokenizer.nextChar()
		tokenizer.CurToken = token.LEFT
	case ")":
		tokenizer.nextChar()
		tokenizer.CurToken = token.RIGHT

	case "+":
		tokenizer.nextChar()
		tokenizer.CurToken = token.ADD
	case "-":
		tokenizer.nextChar()
		tokenizer.CurToken = token.SUB
	case "*":
		tokenizer.nextChar()
		tokenizer.CurToken = token.MUL
	case "/":
		tokenizer.nextChar()
		tokenizer.CurToken = token.DIV

	case "$":
		tokenizer.CurToken = token.END

	default:
		err = fmt.Errorf("can't execute nextToken(): got illegal character at %v", tokenizer.CurPos)
	}

	return tokenizer.CurToken, err
}
