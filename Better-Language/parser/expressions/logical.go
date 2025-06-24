package expressions

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
	"github.com/Chanadu/better-language/utils"
)

type Logical struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (l *Logical) ToGrammarString() string {
	return l.Left.ToGrammarString() + " " + l.Operator.Lexeme + " " + l.Right.ToGrammarString()
}

func (l *Logical) ToReversePolishNotation() string {
	return l.Left.ToReversePolishNotation() + " " + l.Right.ToReversePolishNotation() + " " + l.Operator.Lexeme
}

func (l *Logical) Evaluate(env environment.Environment) (any, error) {
	left, err := l.Left.Evaluate(env)
	if err != nil {
		return nil, err
	}
	var lb, ok bool
	if lb, ok = left.(bool); !ok {
		return nil, utils.CreateRuntimeErrorf(l.Operator.Line, "Left operand must be a boolean")
	}

	switch l.Operator.Type {
	case tokentype.Or:
		if lb {
			return true, nil
		}
	case tokentype.And:
		if !lb {
			return false, nil
		}
	default:
		panic("unhandled default case")
	}

	return l.Right.Evaluate(env)
}
