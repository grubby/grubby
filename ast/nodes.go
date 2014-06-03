package ast

type Node interface{}

type Block struct {
	Statement Node
}

type ConstantInt struct {
	Value int
}
