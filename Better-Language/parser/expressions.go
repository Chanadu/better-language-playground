package parser

import (
	"fmt"

	"github.com/Chanadu/better-language/parser/expressions"
	"github.com/Chanadu/better-language/scanner/tokentype"
)

func (p *parser) parseExpression() expressions.Expression {
	return p.parseAssignment()
}

func (p *parser) parseAssignment() expressions.Expression {
	expr := p.parseTernary()

	if p.match(tokentype.Equal) {
		equals := p.previous()
		value := p.parseAssignment()

		if variable, ok := expr.(*expressions.Variable); ok {
			return &expressions.Assignment{
				Name:  variable.Name,
				Value: value,
			}
		}
		p.err = fmt.Errorf("%s, invalid assignment target", equals.Lexeme)
		return nil
	}

	return expr
}

// Ternary -> Equality ( "?" Expression ":" Expression )?
func (p *parser) parseTernary() expressions.Expression {
	condition := p.parseOr()
	if p.match(tokentype.QuestionMark) {
		trueBranch := p.parseExpression()
		if _, ok := p.consume(tokentype.Colon, "expected ':' after ternary"); ok {
			falseBranch := p.parseExpression()
			return &expressions.Ternary{
				LineNumber:  p.previous().Line,
				Condition:   condition,
				TrueBranch:  trueBranch,
				FalseBranch: falseBranch,
			}
		}
	}
	return condition
}

func (p *parser) parseOr() expressions.Expression {
	expr := p.parseAnd()
	for p.match(tokentype.Or) {
		operator := p.previous()
		right := p.parseAnd()
		expr = &expressions.Logical{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *parser) parseAnd() expressions.Expression {
	expr := p.parseEquality()

	for p.match(tokentype.And) {
		operator := p.previous()
		right := p.parseEquality()
		expr = &expressions.Logical{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

type parseFunc func() expressions.Expression

// LeftAssociativeBinary -> fn ( (tokens) fn )*
func (p *parser) parseLeftAssociativeBinary(fn parseFunc, tokens []tokentype.TokenType) expressions.Expression {
	left := fn()
	for p.match(tokens...) {
		operator := p.previous()
		right := fn()

		left = &expressions.Binary{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

// Equality -> Comparison ( ( "!=" | "==" ) Comparison )*
func (p *parser) parseEquality() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseComparison, []tokentype.TokenType{
		tokentype.NotEqual,
		tokentype.EqualEqual,
	})
}

// Comparison -> Bitwise OR ( ( ">" | ">=" | "<" | "<=" ) Bitwise OR)*
func (p *parser) parseComparison() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseOR, []tokentype.TokenType{
		tokentype.Greater,
		tokentype.GreaterEqual,
		tokentype.Less,
		tokentype.LessEqual,
	})
}

// Bitwise OR -> Bitwise XOR ( ( "|" ) Bitwise XOR)*
func (p *parser) parseBitwiseOR() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseXOR, []tokentype.TokenType{
		tokentype.BitwiseOR,
	})
}

// Bitwise XOR -> Bitwise AND ( ( "^" ) Bitwise AND)*
func (p *parser) parseBitwiseXOR() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseAND, []tokentype.TokenType{
		tokentype.BitwiseXOR,
	})
}

// Bitwise AND -> Bitwise Shift ( ( "&" ) Bitwise Shift)*
func (p *parser) parseBitwiseAND() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseShift, []tokentype.TokenType{
		tokentype.BitwiseAND,
	})
}

// Bitwise Shift -> Term ( ( "<<" | ">>" ) Term )*
func (p *parser) parseBitwiseShift() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseTerm, []tokentype.TokenType{
		tokentype.BitwiseRightShift,
		tokentype.BitwiseLeftShift,
	})
}

// Term -> Factor ( ( "-" | "+" ) Factor )* ;
func (p *parser) parseTerm() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseFactor, []tokentype.TokenType{
		tokentype.Minus,
		tokentype.Plus,
	})
}

// Factor -> Bitwise ( ( "*" | "/" ) Bitwise )* ;
func (p *parser) parseFactor() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseUnary, []tokentype.TokenType{
		tokentype.Star,
		tokentype.Slash,
		tokentype.Percent,
	})
}

// Unary -> ( "-" | "!" | "~" ) Unary | Primary ;
func (p *parser) parseUnary() expressions.Expression {
	if p.match(tokentype.Minus, tokentype.Not, tokentype.BitwiseNOT) {
		operator := p.previous()
		right := p.parseUnary()

		return &expressions.Unary{
			Operator: operator,
			Right:    right,
		}
	}
	return p.parseCall()
}

func (p *parser) parseCall() expressions.Expression {
	expr := p.parsePrimary()

	for {
		if p.match(tokentype.OpeningParentheses) {
			expr = p.finishCall(expr)
		} else {
			break
		}
	}

	return expr
}

func (p *parser) finishCall(callee expressions.Expression) expressions.Expression {
	var args []expressions.Expression
	if !p.check(tokentype.ClosingParentheses) {
		for {
			if len(args) >= 255 {
				p.err = fmt.Errorf("cannot have more than 255 arguments")
			}
			args = append(args, p.parseExpression())
			if !p.match(tokentype.Comma) {
				break
			}
		}
	}

	tt, _ := p.consume(tokentype.ClosingParentheses, "Expect ')' after arguments.")
	return &expressions.Call{
		Callee: callee,
		Paren:  tt,
		Args:   args,
	}
}

// Primary -> Integer | Double | String | True | False | "(" Expression ")" | Null;
func (p *parser) parsePrimary() expressions.Expression {
	if p.match(tokentype.True) {
		return &expressions.Literal{
			Value: true,
		}
	}
	if p.match(tokentype.False) {
		return &expressions.Literal{
			Value: false,
		}
	}
	if p.match(tokentype.Integer, tokentype.Double, tokentype.String) {
		return &expressions.Literal{
			Value: p.previous().Literal,
		}
	}
	if p.match(tokentype.Null) {
		return &expressions.Literal{
			Value: nil,
		}
	}

	if p.match(tokentype.Identifier) {
		return &expressions.Variable{
			Name: p.previous(),
		}
	}

	if p.match(tokentype.OpeningParentheses) {
		expression := p.parseExpression()
		_, _ = p.consume(tokentype.ClosingParentheses, "Expect ')' after expression.")
		return &expressions.Grouping{
			InternalExpression: expression,
		}
	}

	p.err = fmt.Errorf("expect expression, found '%s', line %d", p.peek().Lexeme, p.peek().Line)
	return nil
}
