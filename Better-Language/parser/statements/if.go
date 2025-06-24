package statements

import (
	"errors"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
)

type If struct {
	Condition expressions.Expression
	Then      Statement
	Else      Statement
}

func (i If) Run(env environment.Environment) error {
	condition, err := i.Condition.Evaluate(env)
	if err != nil {
		return err
	}
	var b, ok bool
	if b, ok = condition.(bool); !ok {
		return errors.New("condition must be a boolean")
	}
	if b {
		return i.Then.Run(env)
	} else if i.Else != nil {
		return i.Else.Run(env)
	}

	return nil
}
