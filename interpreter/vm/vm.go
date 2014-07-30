package vm

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/interpreter/vm/builtins"
	"github.com/grubby/grubby/parser"
)

type vm struct {
	ObjectSpace map[string]builtins.Value
}

type VM interface {
	Run(string) (builtins.Value, error)
	Get(string) (builtins.Value, error)
	MustGet(string) builtins.Value
}

func NewVM() VM {
	vm := &vm{
		ObjectSpace: make(map[string]builtins.Value),
	}
	vm.ObjectSpace["Object"] = builtins.NewGlobalObject()
	vm.ObjectSpace["Kernel"] = builtins.NewGlobalKernel()
	return vm
}

func (vm *vm) MustGet(key string) builtins.Value {
	return vm.ObjectSpace[key]
}

func (vm *vm) Get(key string) (builtins.Value, error) {
	val, ok := vm.ObjectSpace[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("NameError: undefined local variable or method '%s' for main:Object"))
	}

	return val, nil
}

func (vm *vm) Run(input string) (builtins.Value, error) {
	lexer := parser.NewLexer(input)
	parser.RubyParse(lexer)

	var val builtins.Value
	for _, statement := range parser.Statements {
		switch statement.(type) {
		case ast.FuncDecl:
			// FIXME: assumes for now this will only ever be at the top level
			funcNode := statement.(ast.FuncDecl)
			method := builtins.NewMethod(funcNode.Name.Name)
			val = method
			vm.ObjectSpace["Kernel"].AddPrivateMethod(method)
		case ast.SimpleString:
			val = builtins.NewString(statement.(ast.SimpleString).Value)
		default:
			panic(fmt.Sprintf("handled unknown statement type: %T:\n\t\n => %#v\n", statement, statement))
		}
	}

	if val == nil {
		return nil, nil
	} else {
		return val, nil
	}
}
