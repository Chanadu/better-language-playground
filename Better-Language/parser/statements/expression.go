package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
)

type Expression struct {
	Expression expressions.Expression
}

func (e *Expression) Run(env environment.Environment) error {
	_, err := e.Expression.Evaluate(env)
	return err
}
