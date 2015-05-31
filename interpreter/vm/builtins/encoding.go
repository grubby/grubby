package builtins

func NewEncodingClass(provider Provider) Class {
	return NewGenericClass("Encoding", "Object", provider)
}
