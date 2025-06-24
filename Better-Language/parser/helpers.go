package parser

import (
	"fmt"
	"strconv"

	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
)

func (p *parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type == tokentype.EndOfFile
}

func (p *parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) check(tokenType tokentype.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *parser) match(tokenTypes ...tokentype.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			_ = p.advance()
			return true
		}
	}
	return false
}

func (p *parser) consume(tokenType tokentype.TokenType, errorMessage string) (t scanner.Token, ok bool) {
	if p.check(tokenType) {

		return p.advance(), true
	}

	token := p.peek()
	location := "EOF"
	if token.Type != tokentype.EndOfFile {
		location = strconv.Itoa(token.Line)
	}

	p.err = fmt.Errorf("%v at %s: %s", token.Lexeme, location, errorMessage)
	return scanner.Token{}, false
}

func (p *parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if _, ok := tokentype.ParseSynchronizationTokens[p.peek().Type]; ok {
			return
		}
		p.advance()
	}
}
