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

type Keyword string

// SQL reserved words
// https://docs.actian.com/psql/psqlv13/index.html#page/sqlref/sqlkword.htm
const (
	AsKeyword     Keyword = "AS"
	AndKeyword    Keyword = "AND"
	ByKeyword     Keyword = "By"
	ColumnKeyword Keyword = "COLUMN"
	CountKeyword  Keyword = "COUNT"
	CreateKeyword Keyword = "CREATE"
	DeleteKeyword Keyword = "DELETE"
	DropKeyword   Keyword = "DROP"
	FalseKeyword  Keyword = "FALSE"
	FromKeyword   Keyword = "FROM"
	GroupKeyword  Keyword = "GROUP"
	IntKeyword    Keyword = "INT"
	LeftKeyword   Keyword = "LEFT"
	LimitKeyword  Keyword = "LIMIT"
	MinKeyword    Keyword = "MIN"
	NullKeyword   Keyword = "NULL"
	OffsetKeyword Keyword = "OFFSET"
	OrKeyword     Keyword = "OR"
	OrderKeyword  Keyword = "ORDER"
	SelectKeyword Keyword = "SELECT"
	SumKeyword    Keyword = "SUM"
	WhereKeyword  Keyword = "WHERE"
)

type TokenKind int

const (
	KeywordKind = iota
	SymbolKind
	IdentifierKind
	StringKind
	NumbericKind
	BoolKind
	NullKind
)

func lexer(stmt string) []Token {

	return []Token{}
}
