package ast

type Node interface{}

type Block struct {
	Statement Node
}

type ConstantInt struct {
	Value int
}

type SimpleString struct {
	Value string
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
