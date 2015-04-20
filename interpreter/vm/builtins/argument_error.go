package builtins

func NewArgumentErrorClass(provider Provider) Class {
	return NewGenericClass("ArgumentError", "Exception", provider)
}
