package builtins

type Value interface {
	String() string
	PrettyPrint() string

	Class() Class

	AddMethod(Method)
	RemoveMethod(Method)
	Method(string) Method
	Methods() []Method

	eigenclassMethods() map[string]Method

	GetInstanceVariable(string) Value
	SetInstanceVariable(string, Value)

	GetClassVariable(string) Value
	SetClassVariable(string, Value)

	IsTruthy() bool
}
