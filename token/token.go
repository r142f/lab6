package token

type Token interface {
	Accept(visitor TokenVisitor)
}

type NumberToken struct {
	Value string
}

func (token *NumberToken) Accept(visitor TokenVisitor) {
	visitor.Visit(token)
}

type Brace int64

const (
	LEFT Brace = iota
	RIGHT
)

func (token Brace) Accept(visitor TokenVisitor) {
	visitor.Visit(token)
}

type Operation int64

const (
	ADD Operation = iota
	SUB
	MUL
	DIV
)

func (token Operation) Accept(visitor TokenVisitor) {
	visitor.Visit(token)
}

type End int64

const (
	END End = iota
)

func (token End) Accept(visitor TokenVisitor) {
	visitor.Visit(token)
}
