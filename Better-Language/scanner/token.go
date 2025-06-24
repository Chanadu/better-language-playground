package scanner

import (
	"github.com/Chanadu/better-language/scanner/tokentype"
)

type Token struct {
	Type    tokentype.TokenType
	Lexeme  string
	Literal any
	Line    int
}
