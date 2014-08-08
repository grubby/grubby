package builtins

type Value interface {
	Methods() []Method
	PrivateMethods() []Method

	AddMethod(Method)
	AddPrivateMethod(Method)

	String() string
	Class() Class
}
