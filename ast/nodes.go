package ast

type Node interface{}

type Block struct {
	Statements []Node
}

type ConstantInt struct {
	Value int
}

type SimpleString struct {
	Value string
}

type Symbol struct {
	Name string
}

type ConstantFloat struct {
	Value float64
}

type BareReference struct {
	Name string
}

type CallExpression struct {
	Func string
	Args []Node
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
}
