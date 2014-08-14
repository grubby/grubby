package builtins

type SymbolValue struct {
	value string
	valueStub
}

func NewSymbol(val string) Value {
	s := &SymbolValue{value: val}
	s.initialize()
	return s
}

func (symbolValue *SymbolValue) String() string {
	return symbolValue.value
}
