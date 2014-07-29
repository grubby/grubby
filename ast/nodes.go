package ast

type Node interface {
}

type Nodes []Node

type Block struct {
	Statements []Node
}

type ConstantInt struct {
	Value int
}

type ConstantFloat struct {
	Value float64
}

type SimpleString struct {
	Value string
}

type Symbol struct {
	Name string
}

type BareReference struct {
	Name string
}

type CallExpression struct {
	Func Node
	Args []Node
}

type FuncDecl struct {
	Name BareReference
	Args []Node
	Body []Node
}

type ClassDecl struct {
	Name       string
	SuperClass Class
	Namespace  string
	Body       []Node
}

type Class struct {
	Name      string
	Namespace string
}

type ModuleDecl struct {
	Name      string
	Namespace string
	Body      []Node
}

type Assignment struct {
	LHS Node
	RHS Node
}

type Boolean struct {
	Value bool
}

type Negation struct {
	Target Node
}

type Complement struct {
	Target Node
}

type Positive struct {
	Target Node
}

type Negative struct {
	Target Node
}

type Addition struct {
	LHS Node
	RHS Node
}

type Subtraction struct {
	LHS Node
	RHS Node
}

type Multiplication struct {
	LHS Node
	RHS Node
}
