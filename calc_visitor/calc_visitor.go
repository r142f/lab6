package calc_visitor

import (
	"fmt"
	"lab6/token"

	"strconv"

	"github.com/golang-collections/collections/stack"
)

type CalcVisitor struct {
	stack stack.Stack
}

func (calcVisitor *CalcVisitor) Calc(tokens []token.Token) (int, error) {
	var err error

	for _, token := range tokens {
		token.Accept(calcVisitor)
	}
	if calcVisitor.stack.Len() != 1 {
		err = fmt.Errorf("can't calculate: after processing stack length != 1")
		return 0, err
	}

	return calcVisitor.stack.Pop().(int), err
}

func (calcVisitor *CalcVisitor) Visit(tkn token.Token) {
	switch t := tkn.(type) {
	case *token.NumberToken:
		number, _ := strconv.Atoi(t.Value)
		calcVisitor.stack.Push(number)

	case token.Operation:
		b, _ := calcVisitor.stack.Pop().(int)
		a, _ := calcVisitor.stack.Pop().(int)

		switch tkn {
		case token.ADD:
			calcVisitor.stack.Push(a + b)
		case token.SUB:
			calcVisitor.stack.Push(a - b)
		case token.MUL:
			calcVisitor.stack.Push(a * b)
		case token.DIV:
			calcVisitor.stack.Push(a / b)
		}
	}
}
