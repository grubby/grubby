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
	Target  Node
	Name    BareReference
	Args    []Node
	Body    []Node
	Rescues []Node
}

func (f FuncDecl) MethodName() string {
	return f.Name.Name
}

func (f FuncDecl) MethodArgs() []MethodParam {
	result := make([]MethodParam, 0, len(f.Args))
	for _, a := range f.Args {
		result = append(result, a.(MethodParam))
	}

	return result
}

type ClassDecl struct {
	Name       string
	SuperClass Class
	Namespace  string
	Body       []Node
}

type Class struct {
	Name              string
	Namespace         string
	IsGlobalNamespace bool
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
type LineNumberConstReference struct{}

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
	Else   []Node
}

type Rescue struct {
	Body      []Node
	Exception RescueException
}

type RescueException struct {
	Var     BareReference
	Classes []Class
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

type Next struct{}
type Redo struct{}
type Break struct{}
type Retry struct{}

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

type Lambda struct {
	Body Block
}

type ProcArg struct {
	Value Node
}

type SwitchStatement struct {
	Condition Node
	Cases     []SwitchCase
	Else      []Node
}

type SwitchCase struct {
	Conditions []Node
	Body       []Node
}

type ConditionalAssignment struct {
	LHS Node
	RHS Node
}

type Range struct {
	Start Node
	End   Node
}

type StarSplat struct {
	Value Node
}

type RescueModifier struct {
	Statement Node
	Rescue    Node
}

type Regex struct {
	Value string
}

type EigenClass struct {
	Target Node
	Body   []Node
}

type Alias struct {
	To   Symbol
	From Symbol
}
