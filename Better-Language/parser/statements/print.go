package statements

import (
	"fmt"
	"syscall/js"

	"github.com/Chanadu/better-language/parser/environment"
	"github.com/Chanadu/better-language/parser/expressions"
)

type Print struct {
	Expression expressions.Expression
}

func (p *Print) Run(env environment.Environment) (err error) {
	v, err := p.Expression.Evaluate(env)
	if err != nil {
		return err
	}
	var output string
	if v == nil {
		output = "null"
		// return nil
	} else {
		output = fmt.Sprintf("%v", v)
	}
	output += "\n"
	js.Global().Call("goPrint", output, false)
	// _, _ = fmt.Println(color.GreenString(output))
	return nil
}
