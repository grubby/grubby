package builtins

type Class interface {
	New() Value
	SuperClass() Class
	String() string
}
