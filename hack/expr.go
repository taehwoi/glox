// generates expression ast
package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var tmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package expressions

import "github.com/taehioum/glox/pkg/token"

type Visitor func(Expr) (any, error)

type Expr interface {
	Accept(Visitor) (any, error)
}
{{range .}}
type {{.Name}} struct {
	{{- range .Fields}}
	{{.}}
	{{- end}}
}

func (expr {{.Name}}) Accept(v Visitor) (any, error) {
	return v(expr)
}
{{end}}`))

func main() {
	defineAst("expressions.go", []string{
		"Assignment : Name token.Token, Value Expr",
		"Binary   : Left Expr, Operator token.Token, Right Expr",
		"Grouping : Expr Expr",
		"Literal  : Value any",
		"Unary    : Operator token.Token, Right Expr",
		"Variable    : Name token.Token",
	})
}

func defineAst(baseName string, types []string) {
	type ast struct {
		Name   string
		Fields []string
	}

	var asts []ast
	for _, t := range types {
		name := strings.TrimSpace(t[:strings.Index(t, ":")])
		fieldStr := strings.TrimSpace(t[strings.Index(t, ":")+1:])

		var fields []string
		for _, f := range strings.Split(fieldStr, ",") {
			fields = append(fields, strings.TrimSpace(f))
		}

		asts = append(asts, ast{
			Name:   name,
			Fields: fields,
		})
	}

	absPath, err := filepath.Abs(baseName)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(absPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmpl.Execute(f, asts)
}
