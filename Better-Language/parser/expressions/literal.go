package expressions

import (
	"fmt"

	"github.com/Chanadu/better-language/parser/environment"
)

type Literal struct {
	Value any
}

func (l *Literal) ToGrammarString() string {
	if l.Value == nil {
		return "null"
	}
	return fmt.Sprint(l.Value)
}

func (l *Literal) ToReversePolishNotation() string {
	if l.Value == nil {
		return "null"
	}
	return fmt.Sprint(l.Value)
}

func (l *Literal) Evaluate(environment.Environment) (any, error) {
	return l.Value, nil
}
