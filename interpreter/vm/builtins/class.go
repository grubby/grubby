package builtins

type Class interface {
	New() Value
	String() string
}
