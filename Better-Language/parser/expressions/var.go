package expressions

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/utils"
)

type Variable struct {
	Name scanner.Token
}

func (v *Variable) ToGrammarString() string {
	// return parenthesizeExpression(u.Operator.Lexeme, u.Right)
	panic("Not implemented")
}

func (v *Variable) ToReversePolishNotation() string {

	panic("Not implemented")
}

func (v *Variable) Evaluate(env environment.Environment) (any, error) {
	val, ok := env.Get(v.Name)
	if !ok {
		return nil, utils.CreateRuntimeErrorf(v.Name.Line, "Undefined variable '%s'", v.Name.Lexeme)
	}
	return val, nil
}
