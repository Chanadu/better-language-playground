package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
	"github.com/Chanadu/better-language/scanner"
)

type Return struct {
	Keyword scanner.Token
	Value   expressions.Expression
}

func (r *Return) Run(env environment.Environment) error {
	if r.Value == nil {
		panic(0)
	}

	value, err := r.Value.Evaluate(env)

	if err != nil {
		return err
	}

	panic(value)
}
