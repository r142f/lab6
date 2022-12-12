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

func (calcVisitor *CalcVisitor) Calc(tokens []token.Token) (result int, err error) {
	if len(tokens) == 0 {
		return
	}

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
		b, ok := calcVisitor.stack.Pop().(int)
		if !ok {
			panic(fmt.Errorf("can't calculate: no operands for operation"))
		}
		a, ok := calcVisitor.stack.Pop().(int)
		if !ok {
			panic(fmt.Errorf("can't calculate: no operands for operation"))
		}

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
