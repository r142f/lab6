package token

type TokenVisitor interface {
	Visit(token Token)
}
