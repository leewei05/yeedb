package yeedb

import (
	"testing"

	"github.com/c2fo/testify/assert"
)

func TestSimpleLexer(t *testing.T) {
	stmt := "SELECT * FROM test;"
	expectTokens := []Token{
		{
			value: "SELECT",
			kind:  IdentifierKind,
			loc: Location{
				line:   0,
				column: 0,
			},
		},
		{
			value: "*",
			kind:  SymbolKind,
			loc: Location{
				line:   0,
				column: 7,
			},
		},
		{
			value: "FROM",
			kind:  KeywordKind,
			loc: Location{
				line:   0,
				column: 9,
			},
		},
		{
			value: "test",
			kind:  IdentifierKind,
			loc: Location{
				line:   0,
				column: 14,
			},
		},
	}
	tokens := lexer(stmt)

	assert.Equal(t, expectTokens, tokens)
}
