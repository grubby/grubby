package builtins

type Value interface {
	Method(string) (Method, error)
	Methods() []Method

	PrivateMethod(string) (Method, error)
	PrivateMethods() []Method

	AddMethod(Method)
	AddPrivateMethod(Method)

	String() string
	Class() Class

	eigenclassMethods() map[string]Method
}
