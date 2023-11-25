package parser

import (
	"fmt"

	"github.com/taehioum/glox/pkg/expressions"
	"github.com/taehioum/glox/pkg/token"
)

type Parser struct {
	tokens []token.Token
	curr   int
}

type InfixParselet interface {
	parse(parser *Parser, left expressions.Expr, token token.Token) (expressions.Expr, error)
	precedence() Precedence
}

type PrefixParselet interface {
	parse(parser *Parser, token token.Token) (expressions.Expr, error)
}

var prefixPraseletsbyTokenType = map[token.Type]PrefixParselet{
	token.PLUS:      UnaryOperatorParselet{},
	token.MINUS:     UnaryOperatorParselet{},
	token.BANG:      UnaryOperatorParselet{},
	token.NUMBER:    LiteralParselet{},
	token.TRUE:      BoolParselet{},
	token.FALSE:     BoolParselet{},
	token.LEFTPAREN: GroupParselet{},
}

var infixPraseletsbyTokenType = map[token.Type]InfixParselet{
	token.PLUS:         TermParselet{},
	token.STAR:         FactorParselet{},
	token.SLASH:        FactorParselet{},
	token.EQUALEQUAL:   EqualityParselet{},
	token.BANGEQUAL:    EqualityParselet{},
	token.LESS:         ComparsionParselet{},
	token.LESSEQUAL:    ComparsionParselet{},
	token.GREATER:      ComparsionParselet{},
	token.GREATEREQUAL: ComparsionParselet{},
}

func Parse(tokens []token.Token) (expressions.Expr, error) {
	parser := Parser{
		tokens: tokens,
		curr:   0,
	}
	expr, err := parser.Parse()
	return expr, err
}

func (p *Parser) Parse() (expressions.Expr, error) {
	expr, err := p.parse(0)
	return expr, err
}

func (p *Parser) parse(precendence Precedence) (expressions.Expr, error) {
	tok := p.consume()

	prefix, ok := prefixPraseletsbyTokenType[tok.Type]
	if !ok {
		return nil, fmt.Errorf("line %d's %s: no parselet for token type %s", tok.Ln, tok.Lexeme, tok.Type)
	}

	left, err := prefix.parse(p, tok)
	if err != nil {
		return left, err
	}

	for precendence < p.precendence() {
		tok := p.consume()

		infix, ok := infixPraseletsbyTokenType[tok.Type]
		if !ok {
			return nil, fmt.Errorf("line %d's %s: no parselet for token type %s", tok.Ln, tok.Lexeme, tok.Type)
		}

		left, err = infix.parse(p, left, tok)
		if err != nil {
			return left, err
		}
	}

	return left, nil
}

func (p *Parser) precendence() Precedence {
	infix, ok := infixPraseletsbyTokenType[p.peek().Type]
	if !ok {
		return 0
	}

	return infix.precedence()
}

// lookahead of distance zero.
func (p *Parser) peek() token.Token {
	return p.tokens[p.curr]
}

func (p *Parser) consume() token.Token {
	tok := p.tokens[p.curr]
	p.curr++
	return tok
}

func (p *Parser) consumeAndCheck(t token.Type, msg string) (token.Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}

	return token.Token{}, fmt.Errorf("line %d's %s: %s", p.peek().Ln, p.peek().Lexeme, msg)
}

func (p *Parser) check(t token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.curr++
	}
	return p.previous()
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.curr-1]
}
