package parser

import (
	"github.com/taehioum/glox/pkg/statements"
	"github.com/taehioum/glox/pkg/token"
)

type StatementParselet struct{}

func (p StatementParselet) parse(parser *Parser, tok token.Token) (statements.Stmt, error) {
	expr, err := parser.parseExpr(0)
	if err != nil {
		return nil, err
	}

	_, err = parser.consumeAndCheck(token.SEMICOLON, "expected ';' after expression")
	if err != nil {
		return nil, err
	}

	switch tok.Type {
	case token.PRINT:
		return statements.Print{Expr: expr}, nil
	default:
		return statements.Expression{Expr: expr}, nil
	}
}