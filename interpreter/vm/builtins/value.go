package builtins

type Value interface {
	String() string
	Class() Class

	AddMethod(Method)
	RemoveMethod(Method)
	Method(string) (Method, error)
	Methods() []Method

	AddPrivateMethod(Method)
	PrivateMethod(string) (Method, error)
	PrivateMethods() []Method

	eigenclassMethods() map[string]Method

	GetInstanceVariable(string) Value
	SetInstanceVariable(string, Value)

	IsTruthy() bool
}
