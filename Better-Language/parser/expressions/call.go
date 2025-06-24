package expressions

import (
	"fmt"

	"github.com/Chanadu/better-language/parser/callable"
	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/scanner"
)

type Call struct {
	Callee Expression
	Paren  scanner.Token
	Args   []Expression
}

func (c *Call) ToGrammarString() string {
	return parenthesizeExpression(c.Paren.Lexeme, c.Args...)
}

func (c *Call) ToReversePolishNotation() string {
	return reversePolishNotation(c.Paren.Lexeme, c.Args...)
}

func (c *Call) Evaluate(env environment.Environment) (any, error) {
	callee, err := c.Callee.Evaluate(env)
	if err != nil {
		return nil, err
	}

	var f callable.Callable
	var ok bool

	if f, ok = callee.(callable.Callable); !ok {
		return nil, fmt.Errorf("can only call functions and classes, %v", c.Callee)
	}

	var args []any
	for _, arg := range c.Args {
		value, err := arg.Evaluate(env)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}

	if len(args) != f.Arity() {
		return nil, fmt.Errorf("expected %d arguments but got %d", f.Arity(), len(args))
	}
	var returnVal any
	err = f.Call(env, args, &returnVal)
	if err != nil {
		return nil, err
	}

	if returnVal == nil {
		return nil, nil
	}

	return returnVal, nil
}
