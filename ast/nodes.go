package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	StringValue() string
}

type Nodes []Node

type Block struct {
	Statements []Node
}

type ConstantInt struct {
	Value int
}

func (i ConstantInt) StringValue() string {
	return fmt.Sprintf("%d", i.Value)
}

type ConstantFloat struct {
	Value float64
}

func (f ConstantFloat) StringValue() string {
	return fmt.Sprintf("%f", f.Value)
}

type SimpleString struct {
	Value string
}

func (s SimpleString) StringValue() string {
	return s.Value
}

type Symbol struct {
	Name string
}

func (s Symbol) StringValue() string {
	return s.Name
}

type BareReference struct {
	Name string
}

func (ref BareReference) StringValue() string {
	return ref.Name
}

type CallExpression struct {
	Func string
	Args []Node
}

func (callexpr CallExpression) StringValue() string {
	args := []string{}
	for _, a := range callexpr.Args {
		args = append(args, a.StringValue())
	}
	return fmt.Sprintf("%s(%s)", callexpr.Func, strings.Join(args, ", "))
}

type FuncDecl struct {
	Name string
	Args []Node
	Body []Node
}

type ClassDefn struct {
	Name       string
	SuperClass string
	Namespace  string
	Body       []Node
}
