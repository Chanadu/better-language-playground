package statements

import (
	"github.com/Chanadu/better-language/parser/environment"
)

type Statement interface {
	Run(env environment.Environment) error
}
