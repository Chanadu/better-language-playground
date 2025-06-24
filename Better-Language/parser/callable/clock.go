package callable

import (
	"time"

	"github.com/Chanadu/better-language/parser/environment"
)

type Clock struct{}

func (c *Clock) Arity() int {
	return 0
}

func (c *Clock) Call(_ environment.Environment, _ []any, res *any) error {
	*res = time.Now().UnixMilli()

	return nil
}

func (c *Clock) String() string {
	return "<clock native callable>"
}
