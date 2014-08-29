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

	loadPath := builtins.NewArrayClass().(builtins.Class).New()
	vm.Globals["$LOAD_PATH"] = loadPath
	vm.Globals["$:"] = loadPath // FIXME: add a test that these are the same object

	objectClass := builtins.NewGlobalObjectClass()
	vm.ObjectSpace["Object"] = objectClass
	vm.ObjectSpace["Kernel"] = builtins.NewGlobalKernelClass()
	vm.ObjectSpace["File"] = builtins.NewFileClass()
	vm.ObjectSpace["ARGV"] = builtins.NewArrayClass().(builtins.Class).New()
	vm.ObjectSpace["Process"] = builtins.NewProcessClass()

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
	main.AddMethod(builtins.NewMethod("puts", func(args ...builtins.Value) (builtins.Value, error) {
		for _, arg := range args {
			os.Stdout.Write([]byte(arg.String() + "\n"))
		}

		return nil, nil
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
		case ast.IfBlock:
			truthy := false
			ifBlock := statement.(ast.IfBlock)
			switch ifBlock.Condition.(type) {
			case ast.Boolean:
				truthy = ifBlock.Condition.(ast.Boolean).Value
			case ast.BareReference:
				truthy = ifBlock.Condition.(ast.BareReference).Name == "nil"
			default:
				truthy = true
			}

			if truthy {
				returnValue, returnErr = vm.executeWithContext(ifBlock.Body, context)
			} else {
				returnValue, returnErr = vm.executeWithContext(ifBlock.Else, context)
			}
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
		case ast.InterpolatedString:
			returnValue = builtins.NewString(statement.(ast.InterpolatedString).Value)
		case ast.GlobalVariable:
			returnValue = vm.Globals[statement.(ast.GlobalVariable).Name]
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

			var target = context
			if callExpr.Target != nil {
				target = vm.ObjectSpace[callExpr.Target.(ast.BareReference).Name]
			}

			if target == nil {
				panic("could not find: " + callExpr.Target.(ast.BareReference).Name)
			}
			method, err := target.Method(callExpr.Func.Name)

			if err != nil {
				err := builtin_errors.NewNameError(callExpr.Func.Name, target.String(), target.Class().String())
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

			switch assignment.LHS.(type) {
			case ast.BareReference:
				ref := assignment.LHS.(ast.BareReference)
				vm.ObjectSpace[ref.Name] = returnValue
			case ast.GlobalVariable:
				globalVar := assignment.LHS.(ast.GlobalVariable)
				vm.Globals[globalVar.Name] = returnValue
			default:
				panic(fmt.Sprintf("unimplemented assignment failure: %#v", assignment.LHS))
			}

		case ast.FileNameConstReference:
			returnValue = builtins.NewString(vm.currentFilename)
		default:
			panic(fmt.Sprintf("handled unknown statement type: %T:\n\t\n => %#v\n", statement, statement))
		}
	}

	return returnValue, returnErr
}
