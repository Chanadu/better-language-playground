package expressions

import (
	"testing"

	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
	"github.com/Chanadu/better-language/utils"
)

func TestGrammarExpression(t *testing.T) {
	e := (&Binary{
		Left: &Unary{
			Operator: scanner.Token{
				Type:    tokentype.Minus,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: &Literal{
				Value: 123,
			},
		},
		Operator: scanner.Token{
			Type:    tokentype.Star,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &Grouping{
			InternalExpression: &Literal{
				Value: 45.67,
			},
		},
	}).ToGrammarString()

	t.Logf("Grammar Expression: %s", e)
	utils.AssertEqual(t, "(* (- 123) (group 45.67))", e)
}

func TestReversePolishNotationExpression(t *testing.T) {
	e := (&Binary{
		Left: &Grouping{
			InternalExpression: &Binary{
				Left: &Literal{
					Value: 1,
				},
				Operator: scanner.Token{
					Type:    tokentype.Plus,
					Lexeme:  "+",
					Literal: nil,
					Line:    1,
				},
				Right: &Literal{
					Value: 2,
				},
			},
		},
		Operator: scanner.Token{
			Type:    tokentype.Star,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &Grouping{
			InternalExpression: &Binary{
				Left: &Literal{
					Value: 4,
				},
				Operator: scanner.Token{
					Type:    tokentype.Minus,
					Lexeme:  "-",
					Literal: nil,
					Line:    1,
				},
				Right: &Literal{
					Value: 3,
				},
			},
		},
	}).ToReversePolishNotation()

	t.Logf("RPN Expression: %s", e)
	utils.AssertEqual(t, "1 2 + 4 3 - *", e)
}
