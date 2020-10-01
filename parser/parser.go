package parser

import (
	"fmt"

	"github.com/ibllex/go-ruler/ast"
	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/token"
	"github.com/ibllex/go-ruler/utils"
)

// Parser generating AST from tokens
type Parser struct {
	lexer        lexer.Lexer
	currentToken token.Token
}

func (p *Parser) eat(t token.Type) error {
	if p.currentToken.Type != t {
		return fmt.Errorf("Invalid syntax, expected %v but %v found", t, p.currentToken.Type)
	}

	p.currentToken = p.lexer.NextToken()
	return nil
}

// expr: term ((AND/OR/XOR) term)*
func (p *Parser) expr() (node ast.AST, err error) {
	if node, err = p.term(); err != nil {
		return
	}

	ops := []interface{}{token.AND, token.OR, token.XOR}
	for utils.InArray(p.currentToken.Type, ops) {
		tk := p.currentToken

		err = p.eat(p.currentToken.Type)
		if err != nil {
			return
		}

		r, err := p.term()
		if err != nil {
			return nil, err
		}

		node = ast.NewLogicalOp(node, tk, r)
	}

	return
}

// term: factor ((EQUAL/GT/LT/NOT_EQUAL/GT_OR_EQUAL/LT_OR_EQUAL) factor)*
func (p *Parser) term() (node ast.AST, err error) {
	node, err = p.factor()
	if err != nil {
		return
	}

	ops := []interface{}{
		token.EQUAL,
		token.GT,
		token.LT,
		token.NOT_EQUAL,
		token.GT_OR_EQUAL,
		token.LT_OR_EQUAL,
	}

	for utils.InArray(p.currentToken.Type, ops) {
		tk := p.currentToken

		err = p.eat(p.currentToken.Type)
		if err != nil {
			return
		}

		r, err := p.factor()
		if err != nil {
			return nil, err
		}

		node = ast.NewLogicalOp(node, tk, r)
	}

	return
}

// factor: INTEGER_CONST | FLOAT_CONST | STRING_CONST | FALSE | TRUE | NULL
// 			functionCall | target | param | LPAREN expr RPAREN
func (p *Parser) factor() (node ast.AST, err error) {
	tk := p.currentToken

	if tk.Type == token.IDENT {
		if p.lexer.CurrentChar() == '(' {
			return p.functionCall()
		}

		return p.target()
	}

	if tk.Type == token.COLON {
		return p.param()
	}

	err = p.eat(tk.Type)
	if err != nil {
		return
	}

	if tk.Type == token.INTEGER_CONST || tk.Type == token.FLOAT_CONST {
		node = ast.NewNum(tk)
		return
	}

	if tk.Type == token.STRING_CONST {
		node = ast.NewStr(tk)
		return
	}

	if tk.Type == token.FALSE || tk.Type == token.TRUE {
		node = ast.NewBoolean(tk)
		return
	}

	if tk.Type == token.NULL {
		node = ast.NewNull()
		return
	}

	if tk.Type == token.LPAREN {
		node, err = p.expr()
		if err != nil {
			return
		}

		err = p.eat(token.RPAREN)
		return
	}

	return nil, fmt.Errorf("Invalid syntax, unexpected %v", tk.Type)
}

// functionCall: IDENT LPAREN (expr (COMMA expr)*)? RPAREN
func (p *Parser) functionCall() (node ast.AST, err error) {
	funcName := p.currentToken.Literal

	if err = p.eat(token.IDENT); err != nil {
		return
	}

	if err = p.eat(token.LPAREN); err != nil {
		return
	}

	params := []ast.AST{}

	if p.currentToken.Type != token.RPAREN {
		param, err := p.expr()
		if err != nil {
			return nil, err
		}

		params = append(params, param)
	}

	for p.currentToken.Type == token.COMMA {
		p.eat(token.COMMA)

		param, err := p.expr()
		if err != nil {
			return nil, err
		}

		params = append(params, param)
	}

	if err = p.eat(token.RPAREN); err != nil {
		return
	}

	node = ast.NewFunctionCall(funcName, params)
	return
}

// target: ident
func (p *Parser) target() (node ast.AST, err error) {
	node, err = p.ident()
	if err != nil {
		return nil, err
	}

	node = ast.NewTarget(node.(ast.Ident))
	return
}

// param: COLON ident
func (p *Parser) param() (node ast.AST, err error) {
	err = p.eat(token.COLON)
	if err != nil {
		return
	}

	node, err = p.ident()
	if err != nil {
		return nil, err
	}

	node = ast.NewParam(node.(ast.Ident))
	return
}

// ident: IDENT (DOT IDENT)*
func (p *Parser) ident() (node ast.AST, err error) {
	path := []string{p.currentToken.Literal}
	err = p.eat(token.IDENT)
	if err != nil {
		return
	}

	for p.currentToken.Type == token.DOT {
		p.eat(token.DOT)
		path = append(path, p.currentToken.Literal)
		err = p.eat(token.IDENT)

		if err != nil {
			return
		}
	}

	node = ast.NewIdent(path)
	return
}

// Parse parser's entry, return top level ast node
func (p *Parser) Parse() (node ast.AST, err error) {
	n, err := p.expr()
	return n, err
}

// New create new Parser by given lexer
func New(lexer lexer.Lexer) Parser {
	p := Parser{lexer: lexer}
	p.currentToken = p.lexer.NextToken()
	return p
}
