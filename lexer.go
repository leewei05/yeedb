package yeedb

type Location struct {
	line   uint
	column uint
}

type Token struct {
	value string
	kind  uint
	loc   Location
}

type TokenKind int

const (
	Keyword = iota
	Symbol
	Identifier
	String
	Numberic
	Bool
	Null
)

func lexer(stmt string) []Token {

	return []Token{}
}
