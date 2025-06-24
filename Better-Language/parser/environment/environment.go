package environment

import (
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/utils"
)

type Environment interface {
	Define(name string, value any) (ok bool)
	Get(name scanner.Token) (value any, ok bool)
	Assign(name scanner.Token, value any) (ok bool)
}

type environment struct {
	enclosing Environment
	values    map[string]any
}

func NewEnvironment(env Environment) Environment {
	return &environment{
		enclosing: env,
		values:    map[string]any{},
	}
}

func (e *environment) Define(name string, value any) (ok bool) {
	_, found := e.values[name]
	if found {
		utils.CreateAndReportParsingErrorf("Variable with name '%s' already defined", name)
		return false
	}
	e.values[name] = value
	return true
}

func (e *environment) Assign(name scanner.Token, value any) (ok bool) {
	_, found := e.values[name.Lexeme]
	if found {
		e.values[name.Lexeme] = value
		return true
	}

	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}

	utils.CreateAndReportParsingErrorf("Undefined variable '%s'", name.Lexeme)
	return false
}

func (e *environment) Get(name scanner.Token) (value any, ok bool) {
	v, found := e.values[name.Lexeme]
	if found {
		return v, true
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	utils.CreateAndReportParsingErrorf("Undefined variable '%s'", name.Lexeme)
	return nil, false
}
