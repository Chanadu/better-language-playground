package scanner

import (
	"github.com/Chanadu/better-language/scanner/tokentype"
)

func (sc *scanner) isAtEnd(offset int) bool {
	return sc.current+offset >= len(sc.source)
}

func (sc *scanner) advanceCurrent() (r rune) {
	r = rune(sc.source[sc.current])
	sc.current++
	return r
}

func (sc *scanner) createToken(tt tokentype.TokenType, literal any) *Token {
	return &Token{
		Type:    tt,
		Lexeme:  sc.source[sc.start:sc.current],
		Literal: literal,
		Line:    sc.lineNumber,
	}
}

func (sc *scanner) match(expected rune) bool {
	if sc.isAtEnd(0) {
		return false
	}
	if rune(sc.source[sc.current]) != expected {
		return false
	}

	sc.current++
	return true
}

func (sc *scanner) peek(offset int) rune {
	if sc.isAtEnd(offset) {
		return rune(0)
	}
	return rune(sc.source[sc.current+offset])
}
