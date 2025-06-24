package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
)

type Block struct {
	Statements []Statement
}

func (b Block) Run(env environment.Environment) error {
	return ExecuteBlock(b.Statements, environment.NewEnvironment(env))
}

func ExecuteBlock(statements []Statement, env environment.Environment) error {
	for _, statement := range statements {
		err := statement.Run(env)
		if err != nil {
			return err
		}
	}
	return nil

}
