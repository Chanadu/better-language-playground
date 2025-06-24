package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
	"github.com/Chanadu/better-language/scanner"
)

type Var struct {
	Name        scanner.Token
	Initializer expressions.Expression
}

func (v *Var) Run(env environment.Environment) error {
	var val any = nil
	var err error
	if v.Initializer != nil {
		val, err = v.Initializer.Evaluate(env)
		if err != nil {
			return err
		}
	}

	env.Define(v.Name.Lexeme, val)
	return nil
}
