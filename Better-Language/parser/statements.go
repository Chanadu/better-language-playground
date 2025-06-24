package parser

import (
	"github.com/Chanadu/better-language/parser/expressions"
	"github.com/Chanadu/better-language/parser/statements"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/scanner/tokentype"
)

func (p *parser) parseDeclaration() (s statements.Statement, ok bool) {
	if p.match(tokentype.Var) {
		s, ok = p.parseVarDeclaration()
	} else if p.match(tokentype.Function) {
		s, ok = p.parseFunction("callable")
	} else {
		s, ok = p.parseStatement()
	}
	if !ok {
		p.synchronize()
	}
	return s, ok
}

func (p *parser) parseVarDeclaration() (s statements.Statement, ok bool) {
	varName, ok := p.consume(tokentype.Identifier, "Expect variable name.")
	if !ok {
		return nil, false
	}

	var initializer expressions.Expression = nil
	if p.match(tokentype.Equal) {
		initializer = p.parseExpression()
	}
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after variable declaration.")
	return &statements.Var{
		Name:        varName,
		Initializer: initializer,
	}, ok
}

func (p *parser) parseStatement() (s statements.Statement, ok bool) {
	if p.match(tokentype.OpeningCurlyBrace) {
		stmts, ok := p.parseBlock()
		if !ok {
			return nil, false
		}
		return statements.Block{
			Statements: stmts,
		}, true
	}
	if p.match(tokentype.Print) {
		return p.parsePrintStatement()
	}
	if p.match(tokentype.If) {
		return p.parseIfStatement()
	}
	if p.match(tokentype.While) {
		return p.parseWhileStatement()
	}
	if p.match(tokentype.For) {
		return p.parseForStatement()
	}
	if p.match(tokentype.Return) {
		return p.parseReturnStatement()
	}

	return p.parseExpressionStatement()
}

func (p *parser) parsePrintStatement() (s statements.Statement, ok bool) {
	expr := p.parseExpression()
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after value.")
	return &statements.Print{
		Expression: expr,
	}, ok
}

func (p *parser) parseExpressionStatement() (s statements.Statement, ok bool) {
	expr := p.parseExpression()
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after expression.")
	return &statements.Expression{
		Expression: expr,
	}, ok
}

func (p *parser) parseBlock() (stmts []statements.Statement, ok bool) {
	stmts = make([]statements.Statement, 0)

	for !p.check(tokentype.ClosingCurlyBrace) && !p.isAtEnd() {
		stmt, ok := p.parseDeclaration()
		if !ok {
			return nil, false
		}
		stmts = append(stmts, stmt)
	}

	_, ok = p.consume(tokentype.ClosingCurlyBrace, "Expect '}' after block.")
	if !ok {
		return nil, false
	}
	p.match(tokentype.Semicolon)
	// _, ok = p.consume(tokentype.Semicolon, "Expect ';' after block('}').")
	// if !ok{
	// 	return nil, false
	// }

	return stmts, true
}

func (p *parser) parseIfStatement() (s statements.Statement, ok bool) {
	_, ok = p.consume(tokentype.OpeningParentheses, "Expect '(' after 'if'.")

	if !ok {
		return nil, false
	}

	expr := p.parseExpression()
	p.consume(tokentype.ClosingParentheses, "Expect ')' after if condition.")
	thenBranch, ok := p.parseStatement()
	if !ok {
		return nil, false
	}
	var elseBranch statements.Statement = nil
	if p.match(tokentype.Else) {
		elseBranch, ok = p.parseStatement()
		if !ok {
			return nil, false
		}
	}
	return &statements.If{
		Condition: expr,
		Then:      thenBranch,
		Else:      elseBranch,
	}, true
}

func (p *parser) parseWhileStatement() (s statements.Statement, ok bool) {
	_, ok = p.consume(tokentype.OpeningParentheses, "Expect '(' after 'while'.")
	if !ok {
		return nil, false
	}

	expr := p.parseExpression()
	_, ok = p.consume(tokentype.ClosingParentheses, "Expect ')' after while condition.")
	if !ok {
		return nil, false
	}

	stmt, ok := p.parseStatement()
	if !ok {
		return nil, false
	}

	return &statements.While{
		Condition: expr,
		Body:      stmt,
	}, true

}

func (p *parser) parseForStatement() (s statements.Statement, ok bool) {
	p.consume(tokentype.OpeningParentheses, "Expect '(' after 'for'.")

	var initializer statements.Statement = nil
	if !p.match(tokentype.Semicolon) {
		if p.match(tokentype.Var) {
			initializer, ok = p.parseVarDeclaration()
			if !ok {
				return nil, false
			}
		} else {
			initializer, ok = p.parseExpressionStatement()
			if !ok {
				return nil, false
			}
		}
	}

	var condition expressions.Expression = nil
	if !p.check(tokentype.Semicolon) {
		condition = p.parseExpression()
	}

	p.consume(tokentype.Semicolon, "Expect ';' after loop condition.")

	var increment expressions.Expression = nil
	if !p.check(tokentype.ClosingParentheses) {
		increment = p.parseExpression()
	}
	p.consume(tokentype.ClosingParentheses, "Expect ')' after for clauses.")

	body, ok := p.parseStatement()
	if !ok {
		return nil, false
	}

	if increment != nil {
		body = &statements.Block{
			Statements: []statements.Statement{
				body,
				&statements.Expression{Expression: increment},
			},
		}
	}
	if condition == nil {
		condition = &expressions.Literal{Value: true}
	}
	body = &statements.While{
		Condition: condition,
		Body:      body,
	}

	if initializer != nil {
		body = &statements.Block{
			Statements: []statements.Statement{
				initializer,
				body,
			},
		}
	}

	return body, true
}

func (p *parser) parseFunction(functionKind string) (s statements.Statement, ok bool) {

	name, ok := p.consume(tokentype.Identifier, "Expect "+functionKind+" name.")
	if !ok {
		return nil, false
	}

	_, ok = p.consume(tokentype.OpeningParentheses, "Expect '(' after "+functionKind+" name.")
	if !ok {
		return nil, false
	}

	var params []scanner.Token
	if !p.check(tokentype.ClosingParentheses) {
		for {
			param, ok := p.consume(tokentype.Identifier, "Expect parameter name.")
			if !ok {
				return nil, false
			}
			params = append(params, param)
			if !p.match(tokentype.Comma) {
				break
			}
		}
	}

	_, ok = p.consume(tokentype.ClosingParentheses, "Expect ')' after parameters.")
	if !ok {
		return nil, false
	}

	_, ok = p.consume(tokentype.OpeningCurlyBrace, "Expect '{' before "+functionKind+" body.")
	if !ok {
		return nil, false
	}

	body, ok := p.parseBlock()
	if !ok {
		return nil, false
	}
	return &statements.Function{
		Name:   name,
		Params: params,
		Body:   body,
	}, true
}

func (p *parser) parseReturnStatement() (s statements.Statement, ok bool) {
	keyword := p.previous()
	var value expressions.Expression = nil
	if !p.check(tokentype.Semicolon) {
		value = p.parseExpression()
	}
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after return value.")
	if !ok {
		return nil, false
	}

	return &statements.Return{
		Keyword: keyword,
		Value:   value,
	}, true

}
