package parser

import (
	"errors"

	"github.com/Chanadu/better-language/globals"
	"github.com/Chanadu/better-language/parser/statements"
	"github.com/Chanadu/better-language/scanner"
)

type Parser interface {
	Parse() ([]statements.Statement, error)
}

type parser struct {
	tokens  []scanner.Token
	current int
	err     error
}

func NewParser(tokenSlice []scanner.Token) Parser {
	return &parser{
		tokens:  tokenSlice,
		current: 0,
	}
}

func (p *parser) Parse() ([]statements.Statement, error) {
	if p.tokens == nil {
		globals.HasErrors = true
		return nil, errors.New("no tokens to parse, need to add tokens to parser")
	}

	// exp := p.parseExpression()
	var stmt []statements.Statement

	for !p.isAtEnd() {
		st, _ := p.parseDeclaration()
		stmt = append(stmt, st)
	}

	if p.err != nil {
		globals.HasErrors = true
	}

	return stmt, p.err
}
