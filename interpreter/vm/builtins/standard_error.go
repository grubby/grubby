package builtins

func NewStandardErrorClass(provider Provider) Class {
	return NewGenericClass("StandardError", "Exception", provider)
}
