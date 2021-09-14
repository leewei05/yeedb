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

func lexer(stmt string) []Token {

	return []Token{}
}
