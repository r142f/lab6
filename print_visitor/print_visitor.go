package print_visitor

import (
	"lab6/token"
	"strings"
)

type PrintVisitor struct {
	sb strings.Builder
}

func (printVisitor *PrintVisitor) Print(tokens []token.Token) string {
	if len(tokens) == 0 {
		return ""
	}

	printVisitor.sb.Reset()

	for _, tkn := range tokens {
		tkn.Accept(printVisitor)
	}

	return printVisitor.sb.String()[:printVisitor.sb.Len()-1]
}

func (printVisitor *PrintVisitor) Visit(tkn token.Token) {
	switch t := tkn.(type) {
	case *token.NumberToken:
		printVisitor.sb.WriteString(t.Value)
		printVisitor.sb.WriteString(" ")

	case token.Operation:
		switch tkn {
		case token.ADD:
			printVisitor.sb.WriteString("+ ")
		case token.SUB:
			printVisitor.sb.WriteString("- ")
		case token.MUL:
			printVisitor.sb.WriteString("* ")
		case token.DIV:
			printVisitor.sb.WriteString("/ ")
		}
	}
}
