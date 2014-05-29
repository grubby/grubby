package ast

type Node interface{}

type Block struct {
	Declaration Node
}

type ConstantInt struct {
	Value int
}
