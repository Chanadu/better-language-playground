package expressions

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/utils"
)

type Assignment struct {
	Name  scanner.Token
	Value Expression
}

func (a Assignment) ToGrammarString() string {
	return parenthesizeExpression("=", a.Value)
}

func (a Assignment) ToReversePolishNotation() string {
	return reversePolishNotation("=", a.Value)
}

func (a Assignment) Evaluate(env environment.Environment) (any, error) {
	val, err := a.Value.Evaluate(env)
	if err != nil {
		return nil, err
	}

	ok := env.Assign(a.Name, val)
	if !ok {
		return nil, utils.CreateRuntimeErrorf(a.Name.Line, "Undefined variable '%s'", a.Name.Lexeme)
	}
	return val, nil
}
