package ast

type Node interface{}

type Nodes []Node

type ConstantInt struct {
	Value int
}

type ConstantFloat struct {
	Value float64
}

type SimpleString struct {
	Value string
}

type InterpolatedString struct {
	Value string
}

type CharacterLiteral struct {
	Value string
}

type Symbol struct {
	Name string
}

type BareReference struct {
	Name string
}

type CallExpression struct {
	Target Node
	Func   BareReference
	Args   []Node
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

type Array struct {
	Nodes []Node
}

type Hash struct {
	Pairs []HashKeyValuePair
}

type HashKeyValuePair struct {
	Key   Node
	Value Node
}

type GlobalVariable struct {
	Name string
}

type InstanceVariable struct {
	Name string
}

type ClassVariable struct {
	Name string
}

type FileNameConstReference struct{}

type Block struct {
	Args []Node
	Body []Node
}

type IfBlock struct {
	Condition Node
	Body      []Node
	Else      []Node
}

type Subshell struct {
	Command string
}

type Group struct {
	Body []Node
}

type Begin struct {
	Body   []Node
	Rescue []Node
}

type Rescue struct {
	Body      []Node
	Exception RescueException
}

type RescueException struct {
	Var   BareReference
	Class BareReference
}

type MethodParam struct {
	Name         BareReference
	DefaultValue Node
	IsSplat      bool
	IsProc       bool
}

type Ternary struct {
	Condition Node
	True      Node
	False     Node
}

type Yield struct {
	Value Node
}

type Return struct {
	Value Node
}

type Break struct{}
type Next struct{}

type Loop struct {
	Condition Node
	Body      []Node
}

type WeakLogicalAnd struct {
	LHS Node
	RHS Node
}

type WeakLogicalOr struct {
	LHS Node
	RHS Node
}
