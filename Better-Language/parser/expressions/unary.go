package expressions

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
	"github.com/Chanadu/better-language/utils"
)

type Unary struct {
	Operator scanner.Token
	Right    Expression
}

func (u *Unary) ToGrammarString() string {
	return parenthesizeExpression(u.Operator.Lexeme, u.Right)
}

func (u *Unary) ToReversePolishNotation() string {
	return reversePolishNotation(u.Operator.Lexeme, u.Right)
}

func (u *Unary) Evaluate(env environment.Environment) (any, error) {
	right, _ := u.Right.Evaluate(env)
	switch u.Operator.Type {
	case tokentype.Minus:
		d, dOk := right.(float64)
		i, iOk := right.(int64)

		if !dOk && !iOk {
			return nil, utils.CreateRuntimeErrorf(u.Operator.Line, "expect a number(double or int) after '-'")
		}
		if dOk && iOk {
			panic("Number is both int and float64")
		}

		if dOk {
			return -d, nil
		}

		return -i, nil
	case tokentype.Not:
		b, ok := right.(bool)
		if !ok {
			return nil, utils.CreateRuntimeErrorf(u.Operator.Line, "expect a boolean after '!'")
		}
		return !b, nil

	case tokentype.BitwiseNOT:
		i, ok := right.(int64)
		if !ok {
			return nil, utils.CreateRuntimeErrorf(u.Operator.Line, "expect an integer after '~'")
		}
		return ^i, nil
	default:
		panic("Unknown unary operator")
	}
}
