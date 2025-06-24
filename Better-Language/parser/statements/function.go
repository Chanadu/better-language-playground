package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/utils"
)

type CallableFunction struct {
	Declaration *Function
}

func (f *CallableFunction) Call(global environment.Environment, args []any, res *any) error {
	env := environment.NewEnvironment(global)
	for i, param := range f.Declaration.Params {
		env.Define(param.Lexeme, args[i])
	}

	defer func() {
		if r := recover(); r != nil {
			*res = r
		}
	}()

	err := ExecuteBlock(f.Declaration.Body, env)
	if err != nil {
		return err
	}
	return nil
}

func (f *CallableFunction) Arity() int {
	return len(f.Declaration.Params)
}

func (f *CallableFunction) String() string {
	return "<fn " + f.Declaration.Name.Lexeme + ">"
}

type Function struct {
	Name   scanner.Token
	Params []scanner.Token
	Body   []Statement
}

func (f *Function) Run(env environment.Environment) error {
	cf := &CallableFunction{Declaration: f}

	ok := env.Define(f.Name.Lexeme, cf)
	if !ok {
		return utils.CreateRuntimeErrorf(f.Name.Line, "function %s already defined", f.Name.Lexeme)
	}

	return nil
}
