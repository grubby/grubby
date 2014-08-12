package builtins

type SymbolValue struct {
	value string
	valueStub
}

func NewSymbol(val string) Value {
	return &SymbolValue{value: val}
}

func (symbolValue *SymbolValue) String() string {
	return symbolValue.value
}
