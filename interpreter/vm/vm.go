package vm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/interpreter/vm/builtins"
	builtin_errors "github.com/grubby/grubby/interpreter/vm/builtins/errors"
	"github.com/grubby/grubby/parser"
)

type vm struct {
	currentFilename string
	ObjectSpace     map[string]builtins.Value
	Globals         map[string]builtins.Value
	KnownSymbols    map[string]builtins.Value
}

type VM interface {
	Run(string) (builtins.Value, error)
	Get(string) (builtins.Value, error)
	MustGet(string) builtins.Value

	Set(string, builtins.Value)

	Symbols() map[string]builtins.Value
}

func NewVM(name string) VM {
	vm := &vm{
		currentFilename: name,
		Globals:         make(map[string]builtins.Value),
		ObjectSpace:     make(map[string]builtins.Value),
		KnownSymbols:    make(map[string]builtins.Value),
	}

	loadPath := builtins.NewArrayClass()
	vm.Globals["$LOAD_PATH"] = loadPath
	vm.Globals["$:"] = loadPath // FIXME: add a test that these are the same object

	objectClass := builtins.NewGlobalObjectClass()
	vm.ObjectSpace["Object"] = objectClass
	vm.ObjectSpace["Kernel"] = builtins.NewGlobalKernelClass()
	vm.ObjectSpace["File"] = builtins.NewFileClass()

	main := objectClass.(builtins.Class).New()
	main.AddMethod(builtins.NewMethod("to_s", func(args ...builtins.Value) (builtins.Value, error) {
		return builtins.NewString("main"), nil
	}))
	main.AddMethod(builtins.NewMethod("require", func(args ...builtins.Value) (builtins.Value, error) {
		fileName := args[0].(*builtins.StringValue).String()
		fileName = fileName[1 : len(fileName)-1]

		for _, pathStr := range loadPath.(*builtins.Array).Members() {
			path := pathStr.(*builtins.StringValue)
			fullPath := filepath.Join(path.String(), fileName+".rb")
			file, err := os.Open(fullPath)
			if err != nil {
				continue
			}

			contents, err := ioutil.ReadAll(file)

			if err == nil {
				originalName := vm.currentFilename
				defer func() {
					vm.currentFilename = originalName
				}()

				vm.currentFilename = file.Name()
				vm.Run(string(contents))
				return nil, nil
			}
		}

		err := fmt.Sprintf("LoadError: cannot load such file -- %s", fileName)
		return nil, builtin_errors.NewLoadError(err)

	}))
	vm.ObjectSpace["main"] = main

	return vm
}

func (vm *vm) MustGet(key string) builtins.Value {
	val, ok := vm.ObjectSpace[key]
	if ok {
		return val
	}

	val, ok = vm.Globals[key]
	if ok {
		return val
	}

	return nil
}

func (vm *vm) Get(key string) (builtins.Value, error) {
	val, ok := vm.ObjectSpace[key]
	if ok {
		return val, nil
	}

	val, ok = vm.Globals[key]
	if ok {
		return val, nil
	}

	return nil, errors.New(fmt.Sprintf("'%s' is undefined", key))
}

func (vm *vm) Set(key string, value builtins.Value) {
	vm.ObjectSpace[key] = value
}

func (vm *vm) Symbols() map[string]builtins.Value {
	return vm.KnownSymbols
}

func (vm *vm) Run(input string) (builtins.Value, error) {
	lexer := parser.NewLexer(input)
	parser.RubyParse(lexer)

	main := vm.ObjectSpace["main"]
	return vm.executeWithContext(parser.Statements, main)
}

func (vm *vm) executeWithContext(statements []ast.Node, context builtins.Value) (builtins.Value, error) {
	var (
		returnValue builtins.Value
		returnErr   error
	)
	for _, statement := range statements {
		switch statement.(type) {
		case ast.FuncDecl:
			// FIXME: assumes for now this will only ever be at the top level
			// it seems like this should be replaced with context, but that's really context for calling methods, not
			// necessarily for defining new methods
			funcNode := statement.(ast.FuncDecl)
			method := builtins.NewMethod(funcNode.Name.Name, func(args ...builtins.Value) (builtins.Value, error) {
				return nil, nil
			})
			returnValue = method
			vm.ObjectSpace["Kernel"].AddPrivateMethod(method)
		case ast.SimpleString:
			returnValue = builtins.NewString(statement.(ast.SimpleString).Value)
		case ast.ConstantInt:
			returnValue = builtins.NewInt(statement.(ast.ConstantInt).Value)
		case ast.ConstantFloat:
			returnValue = builtins.NewFloat(statement.(ast.ConstantFloat).Value)
		case ast.Symbol:
			name := statement.(ast.Symbol).Name
			maybe, ok := vm.KnownSymbols[name]
			if !ok {
				returnValue = builtins.NewSymbol(name)
				vm.KnownSymbols[name] = returnValue
			} else {
				returnValue = maybe
			}
		case ast.BareReference:
			name := statement.(ast.BareReference).Name
			maybe, ok := vm.ObjectSpace[name]
			if ok {
				returnValue = maybe
			} else {
				returnErr = builtin_errors.NewNameError(name, context.String(), context.Class().String())
			}
		case ast.CallExpression:
			var method builtins.Method
			callExpr := statement.(ast.CallExpression)
			method, err := context.Method(callExpr.Func.Name)

			if err != nil {
				err := builtin_errors.NewNameError(callExpr.Func.Name, context.String(), context.Class().String())
				return nil, err
			}

			args := []builtins.Value{}
			for _, astArgument := range callExpr.Args {
				arg, err := vm.executeWithContext([]ast.Node{astArgument}, context)
				if err != nil {
					return nil, err
				}

				args = append(args, arg)
			}

			returnValue, returnErr = method.Execute(args...)
		case ast.Assignment:
			assignment := statement.(ast.Assignment)
			returnValue, err := vm.executeWithContext([]ast.Node{assignment.RHS}, context)
			if err != nil {
				return nil, err
			}

			ref, ok := assignment.LHS.(ast.BareReference)
			if !ok {
				panic("unimplemented")
			}

			vm.ObjectSpace[ref.Name] = returnValue

		case ast.FileNameConstReference:
			returnValue = builtins.NewString(vm.currentFilename)
		default:
			panic(fmt.Sprintf("handled unknown statement type: %T:\n\t\n => %#v\n", statement, statement))
		}
	}

	return returnValue, returnErr
}
