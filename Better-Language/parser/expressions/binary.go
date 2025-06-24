package expressions

import (
	"fmt"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
	"github.com/Chanadu/better-language/utils"
)

type Binary struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (b *Binary) ToGrammarString() string {
	return parenthesizeExpression(b.Operator.Lexeme, b.Left, b.Right)
}

func (b *Binary) ToReversePolishNotation() string {
	return reversePolishNotation(b.Operator.Lexeme, b.Left, b.Right)
}

func (b *Binary) Evaluate(env environment.Environment) (any, error) {
	left, err := b.Left.Evaluate(env)
	if err != nil {
		return nil, err
	}
	right, err := b.Right.Evaluate(env)
	if err != nil {
		return nil, err
	}

	switch b.Operator.Type {
	case tokentype.NotEqual:
		return left != right, nil
	case tokentype.EqualEqual:
		return left == right, nil
	case tokentype.Greater:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf > rf, nil
		}
		if liOk && riOk {
			return li > ri, nil
		}
		if lfOk {
			return lf > float64(ri), nil
		}
		return float64(li) > rf, nil
	case tokentype.GreaterEqual:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf >= rf, nil
		}
		if liOk && riOk {
			return li >= ri, nil
		}
		if lfOk {
			return lf >= float64(ri), nil
		}
		return float64(li) >= rf, nil
	case tokentype.Less:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf < rf, nil
		}
		if liOk && riOk {
			return li < ri, nil
		}
		if lfOk {
			return lf < float64(ri), nil
		}
		return float64(li) < rf, nil
	case tokentype.LessEqual:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateErrorf("Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf <= rf, nil
		}
		if liOk && riOk {
			return li <= ri, nil
		}
		if lfOk {
			return lf <= float64(ri), nil
		}
		return float64(li) <= rf, nil
	case tokentype.BitwiseOR:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return li | ri, nil
	case tokentype.BitwiseXOR:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return li ^ ri, nil
	case tokentype.BitwiseAND:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return li & ri, nil
	case tokentype.BitwiseLeftShift:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		if li < 0 {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be positive", b.Operator.Lexeme)
		}
		if ri < 0 {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be positive", b.Operator.Lexeme)
		}
		return li << uint(ri), nil
	case tokentype.BitwiseRightShift:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		if li < 0 {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be positive", b.Operator.Lexeme)
		}
		if ri < 0 {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be positive", b.Operator.Lexeme)
		}
		return li >> uint(ri), nil
	case tokentype.Minus:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf - rf, nil
		}
		if liOk && riOk {
			return li - ri, nil
		}
		if lfOk {
			return lf - float64(ri), nil
		}
		return float64(li) - rf, nil
	case tokentype.Plus:
		if _, ok := left.(string); ok {
			if _, ok := right.(string); ok {
				return left.(string) + right.(string), nil
			}
			rs := fmt.Sprintf("%#v", right)
			return left.(string) + rs, nil

			// return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be string", b.Operator.Lexeme)
		}
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf + rf, nil
		}
		if liOk && riOk {
			return li + ri, nil
		}
		if lfOk {
			return lf + float64(ri), nil
		}
		return float64(li) + rf, nil
	case tokentype.Star:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}

		if lfOk && rfOk {
			return lf * rf, nil
		}
		if liOk && riOk {
			return li * ri, nil
		}
		if lfOk {
			return lf * float64(ri), nil
		}
		return float64(li) * rf, nil
	case tokentype.Slash:
		lf, lfOk := left.(float64)
		rf, rfOk := right.(float64)
		li, liOk := left.(int64)
		ri, riOk := right.(int64)

		if !lfOk && !liOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Left Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if !rfOk && !riOk {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Right Operand of (%s) must be numbers", b.Operator.Lexeme)
		}
		if rf == 0 || ri == 0 {
			return nil, utils.CreateRuntimeErrorf(b.Operator.Line, "Division by zero")
		}
		if lfOk && rfOk {
			return lf / rf, nil
		}
		if liOk && riOk {
			return li / ri, nil
		}
		if lfOk {
			return lf / float64(ri), nil
		}
		return float64(li) / rf, nil
	case tokentype.Percent:
		li, ri, err := integerBinaryExpression(b.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return li % ri, nil
	default:
		panic("Unknown binary operator")
	}
}

func integerBinaryExpression(operator scanner.Token, left any, right any) (l int64, r int64, err error) {
	l, lOk := left.(int64)
	r, rOk := right.(int64)
	if !lOk && !rOk {
		return 0, 0, utils.CreateRuntimeErrorf(operator.Line, "Left and Right Operand of (%s) must be int", operator.Lexeme)
	}
	if !lOk {
		return 0, 0, utils.CreateRuntimeErrorf(operator.Line, "Left Operand of (%s) must be int", operator.Lexeme)
	}
	if !rOk {
		return 0, 0, utils.CreateRuntimeErrorf(operator.Line, "Right Operand of (%s) must be int", operator.Lexeme)
	}

	return l, r, nil
}
