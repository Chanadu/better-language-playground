package statements

import (
	"errors"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
)

type While struct {
	Condition expressions.Expression
	Body      Statement
}

func (w While) Run(env environment.Environment) error {

	for {
		c, err := w.Condition.Evaluate(env)
		if err != nil {
			return err
		}

		var cb, ok bool
		if cb, ok = c.(bool); !ok {
			return errors.New("condition must be a boolean")
		}
		if !cb {
			break
		}

		err = w.Body.Run(env)
		if err != nil {
			return err
		}
	}

	return nil
}
