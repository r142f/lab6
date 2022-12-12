package parser_visitor

import (
	"fmt"
	"lab6/token"
	"lab6/tokenizer"

	"github.com/golang-collections/collections/stack"
)

type ParserVisitor struct {
	tokens    []token.Token
	stack     stack.Stack
	tokenizer *tokenizer.Tokenizer
}

func (parserVisitor *ParserVisitor) Parse(input string) (tokens []token.Token, err error) {
	defer func() {
		parserVisitor.tokens = []token.Token{}

		switch p := recover(); p.(type) {
		case error:
			err = p.(error)
		default:
			if p != nil {
				panic(p)
			}
		}
	}()

	parserVisitor.tokenizer = &tokenizer.Tokenizer{Input: input}
	parserVisitor.tokenizer.Init()

	var tkn token.Token
	for tkn, err = parserVisitor.tokenizer.NextToken(); tkn != token.END && err == nil; tkn, err = parserVisitor.tokenizer.NextToken() {
		tkn.Accept(parserVisitor)
	}

	if err != nil {
		return nil, err
	}

	for parserVisitor.stack.Len() > 0 {
		tkn := parserVisitor.stack.Pop().(token.Token)

		switch tkn.(type) {
		case token.Operation:
			parserVisitor.tokens = append(parserVisitor.tokens, tkn)

		default:
			err = fmt.Errorf("can't parse: parentheses in an expression are inconsistent")
		}
	}

	return parserVisitor.tokens, err
}

func hasGEPriority(op1 token.Operation, op2 token.Token) bool {
	return op2 == token.MUL ||
		op2 == token.DIV ||
		(op2 == token.ADD ||
			op2 == token.SUB) &&
			(op1 == token.ADD ||
				op1 == token.SUB)
}

func (parserVisitor *ParserVisitor) Visit(tkn token.Token) {
	switch tknOp := tkn.(type) {
	case *token.NumberToken:
		parserVisitor.tokens = append(parserVisitor.tokens, tkn)

	case token.Brace:
		if tkn == token.LEFT {
			parserVisitor.stack.Push(tkn)
		} else {
			for parserVisitor.stack.Len() > 0 && parserVisitor.stack.Peek() != token.LEFT {
				parserVisitor.tokens = append(parserVisitor.tokens, parserVisitor.stack.Pop().(token.Token))
			}

			if parserVisitor.stack.Len() == 0 {
				panic(fmt.Errorf("can't parse: didn't get expected left brace"))
			}

			parserVisitor.stack.Pop()
		}

	case token.Operation:
		for parserVisitor.stack.Len() > 0 && hasGEPriority(tknOp, parserVisitor.stack.Peek().(token.Token)) {
			parserVisitor.tokens = append(parserVisitor.tokens, parserVisitor.stack.Pop().(token.Token))
		}

		parserVisitor.stack.Push(tkn)
	}
}
