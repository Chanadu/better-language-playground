package callable

import (
	"github.com/Chanadu/better-language/parser/environment"
)

type Callable interface {
	Call(environment.Environment, []any, *any) error
	Arity() int
	String() string
}
