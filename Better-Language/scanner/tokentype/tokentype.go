package tokentype

type TokenType int

const (
	// Base Nil Character
	Base TokenType = iota

	// Single Character Token

	OpeningParentheses
	ClosingParentheses
	OpeningCurlyBrace
	ClosingCurlyBrace

	Comma
	Dot
	Semicolon

	Minus
	Plus
	Star
	Percent
	Slash

	// One or Two Character Token

	NotEqual
	Not
	EqualEqual
	Equal

	GreaterEqual
	Greater

	LessEqual
	Less

	Or
	And

	BitwiseLeftShift
	BitwiseRightShift
	BitwiseAND
	BitwiseOR
	BitwiseXOR
	BitwiseNOT

	// Literals

	Identifier
	String
	Integer
	Double

	// Keywords

	False
	True

	If
	Else
	For
	While

	Break
	Return
	Continue

	Function

	Print
	Var

	Null

	EndOfFile

	QuestionMark
	Colon

	// Won't use this stuff

	// Class
	// This
	// Super
)

//go:generate stringer -type=TokenType

var KeywordsToTokenType = map[string]TokenType{
	"false":    False,
	"true":     True,
	"if":       If,
	"else":     Else,
	"for":      For,
	"while":    While,
	"break":    Break,
	"return":   Return,
	"continue": Continue,
	"function": Function,
	"print":    Print,
	"var":      Var,
	"null":     Null,
	// "class":    Class,
	// "this":     This,
}

var NewLineSemicolonTokens = map[TokenType]struct{}{
	ClosingParentheses: {},
	ClosingCurlyBrace:  {},
	Identifier:         {},
	String:             {},
	Integer:            {},
	Double:             {},
	True:               {},
	False:              {},
	Break:              {},
	Return:             {},
	Continue:           {},
	Null:               {},
}

var ParseSynchronizationTokens = map[TokenType]struct{}{
	Function: {},
	Var:      {},
	If:       {},
	While:    {},
	For:      {},
	Return:   {},
	Break:    {},
	Continue: {},
	Print:    {},
}
