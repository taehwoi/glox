package expressions

import (
	"fmt"

	"github.com/taehioum/glox/pkg/token"
)

type Visitor func(Expr) (any, error)

type Expr interface {
	Accept(Visitor) (any, error)
}

type Assignment struct {
	Name  token.Token
	Value Expr
}

func (expr Assignment) Accept(v Visitor) (any, error) {
	return v(expr)
}

func (expr Assignment) String() string {
	return fmt.Sprintf("Assignment{Name: %s, Value: %s}", expr.Name, expr.Value)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr Binary) Accept(v Visitor) (any, error) {
	return v(expr)
}

type Grouping struct {
	Expr Expr
}

func (expr Grouping) Accept(v Visitor) (any, error) {
	return v(expr)
}

type Literal struct {
	Value any
}

func (expr Literal) Accept(v Visitor) (any, error) {
	return v(expr)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (expr Unary) Accept(v Visitor) (any, error) {
	return v(expr)
}

type Variable struct {
	Name token.Token
}

func (expr Variable) Accept(v Visitor) (any, error) {
	return v(expr)
}

type Logical struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr Logical) Accept(v Visitor) (any, error) {
	return v(expr)
}
