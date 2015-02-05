package vm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

type vm struct {
	currentFilename string

	stack          *CallStack
	ObjectSpace    map[string]Value
	CurrentGlobals map[string]Value
	CurrentSymbols map[string]Value
	CurrentClasses map[string]Class
	CurrentModules map[string]Module
	singletons     map[string]Value

	localVariableStack *localVariableStack
}

type VM interface {
	Run(string) (Value, error)
	Get(string) (Value, error)
	MustGet(string) Value

	GetClass(string) (Class, error)
	MustGetClass(string) Class

	Set(string, Value)

	Symbols() map[string]Value
	Globals() map[string]Value
	Classes() map[string]Class
	Modules() map[string]Module

	ClassProvider
	SingletonProvider
}

func NewVM(rubyHome, name string) VM {
	vm := &vm{
		currentFilename:    name,
		stack:              NewCallStack(),
		CurrentGlobals:     make(map[string]Value),
		ObjectSpace:        make(map[string]Value),
		CurrentSymbols:     make(map[string]Value),
		CurrentModules:     make(map[string]Module),
		localVariableStack: newLocalVariableStack(),
		singletons:         make(map[string]Value),
	}
	vm.registerBuiltinClassesAndModules()

	loadPath, _ := vm.CurrentClasses["Array"].New(vm, vm)
	loadPath.(*Array).Append(NewString(filepath.Join(rubyHome, "lib"), vm, vm))

	vm.CurrentGlobals["LOAD_PATH"] = loadPath
	vm.CurrentGlobals[":"] = loadPath
	vm.ObjectSpace["ARGV"], _ = vm.CurrentClasses["Array"].New(vm, vm)

	main, _ := vm.CurrentClasses["Object"].New(vm, vm)
	main.AddMethod(NewNativeMethod("to_s", vm, vm, func(self Value, block Block, args ...Value) (Value, error) {
		return NewString("main", vm, vm), nil
	}))
	vm.ObjectSpace["main"] = main

	return vm
}

func (vm *vm) registerBuiltinClassesAndModules() {
	vm.CurrentClasses = map[string]Class{}
	vm.CurrentModules = map[string]Module{}

	basicObjectClass := NewBasicObjectClass(vm)
	vm.CurrentClasses["BasicObject"] = basicObjectClass

	objectClass := NewGlobalObjectClass(vm, vm)
	vm.CurrentClasses["Object"] = objectClass

	classClass := NewClassClass(vm, vm)
	vm.CurrentClasses["Class"] = classClass

	moduleClass := NewModuleClass(vm, vm)
	vm.CurrentClasses["Module"] = moduleClass
	vm.CurrentModules["Comparable"] = NewComparableModule(vm, vm)
	vm.CurrentModules["Kernel"] = NewGlobalKernelModule(vm, vm)
	vm.CurrentModules["Process"] = NewProcessModule(vm)

	// FIXME: this should be private, but method resolution fails
	vm.CurrentModules["Kernel"].AddMethod(NewNativeMethod("require", vm, vm, func(self Value, block Block, args ...Value) (Value, error) {
		fileName := args[0].(*StringValue).RawString()
		if fileName == "rubygems" {
			// don't "require 'rubygems'"
			return vm.singletons["false"], nil
		}

		loadPath := vm.CurrentGlobals["LOAD_PATH"]
		for _, pathStr := range loadPath.(*Array).Members() {
			path := pathStr.(*StringValue)
			fullPath := filepath.Join(path.RawString(), fileName+".rb")
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
				_, rubyErr := vm.Run(string(contents))
				if rubyErr != nil {
					return nil, rubyErr
				}

				return vm.singletons["true"], nil
			}
		}

		errorMessage := fmt.Sprintf("LoadError: cannot load such file -- %s", fileName)
		return nil, NewLoadError(errorMessage, vm.stack.String())
	}))

	/* BEGIN RUNTIME TRICKERY
	There's a cycle in ruby's builtin object graph
	There are classes that refer to each other (Module, Class)
	As well as including the Kernel module
	*/
	objectClass.Include(vm.CurrentModules["Kernel"])
	moduleClass.Include(vm.CurrentModules["Kernel"])
	classClass.(*ClassValue).SetSuperClass()
	objectClass.(*ObjectClass).SetSuperClass()
	basicObjectClass.(*BasicObjectClass).SetSuperClass()
	// END RUNTIME TRICKERY

	vm.CurrentClasses["IO"] = NewIOClass(vm)
	vm.CurrentClasses["Array"] = NewArrayClass(vm, vm)
	vm.CurrentClasses["Hash"] = NewHashClass(vm, vm)
	vm.CurrentClasses["TrueClass"] = NewTrueClass(vm)
	vm.CurrentClasses["File"] = NewFileClass(vm, vm)
	vm.CurrentClasses["FalseClass"] = NewFalseClass(vm)
	vm.CurrentClasses["NilClass"] = NewNilClass(vm)
	vm.CurrentClasses["String"] = NewStringClass(vm, vm)
	vm.CurrentClasses["Numeric"] = NewNumericClass(vm)
	vm.CurrentClasses["Integer"] = NewIntegerClass(vm)
	vm.CurrentClasses["Fixnum"] = NewFixnumClass(vm, vm)
	vm.CurrentClasses["Float"] = NewFloatClass(vm)
	vm.CurrentClasses["Symbol"] = NewSymbolClass(vm)

	vm.singletons["nil"], _ = vm.CurrentClasses["NilClass"].New(vm, vm)
	vm.singletons["true"], _ = vm.CurrentClasses["TrueClass"].New(vm, vm)
	vm.singletons["false"], _ = vm.CurrentClasses["FalseClass"].New(vm, vm)
}

func (vm *vm) MustGet(key string) Value {
	val, err := vm.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func (vm *vm) Get(key string) (Value, error) {
	val, ok := vm.ObjectSpace[key]
	if ok {
		return val, nil
	}

	val, ok = vm.CurrentGlobals[key]
	if ok {
		return val, nil
	}

	class, ok := vm.CurrentClasses[key]
	if ok {
		return class, nil
	}

	module, ok := vm.CurrentModules[key]
	if ok {
		return module, nil
	}

	return nil, errors.New(fmt.Sprintf("'%s' is undefined", key))
}

func (vm *vm) GetClass(name string) (Class, error) {
	for keyName, class := range vm.CurrentClasses {
		if keyName == name {
			return class, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Class '%s' not found", name))
}

func (vm *vm) MustGetClass(name string) Class {
	for keyName, class := range vm.CurrentClasses {
		if keyName == name {
			return class
		}
	}

	panic(fmt.Sprintf("class '%s' requested, but does not exist", name))
}

func (vm *vm) Set(key string, value Value) {
	vm.ObjectSpace[key] = value
}

func (vm *vm) Symbols() map[string]Value {
	return vm.CurrentSymbols
}

func (vm *vm) Globals() map[string]Value {
	return vm.CurrentGlobals
}

func (vm *vm) Classes() map[string]Class {
	return vm.CurrentClasses
}

func (vm *vm) Modules() map[string]Module {
	return vm.CurrentModules
}

type ParseError struct {
	Filename string
}

func NewParseError(filename string) *ParseError {
	return &ParseError{Filename: filename}
}

func (err *ParseError) Error() string {
	return "parse error"
}

func (vm *vm) Run(input string) (Value, error) {
	parser.Statements = []ast.Node{}
	lexer := parser.NewLexer(input)
	result := parser.RubyParse(lexer)
	if result != 0 {
		return nil, NewParseError(vm.currentFilename)
	}

	main := vm.ObjectSpace["main"]
	vm.stack.Unshift("main", vm.currentFilename)
	defer vm.stack.Shift()

	vm.localVariableStack.unshift()
	defer vm.localVariableStack.shift()
	return vm.executeWithContext(main, parser.Statements...)
}

func (vm *vm) executeWithContext(context Value, statements ...ast.Node) (Value, error) {
	var (
		returnValue Value
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
				returnValue, returnErr = vm.executeWithContext(context, ifBlock.Body...)
			} else {
				returnValue, returnErr = vm.executeWithContext(context, ifBlock.Else...)
			}
		case ast.Alias:
			// FIXME: assumes that the context will be a module, but could also be a class
			aliasNode := statement.(ast.Alias)
			contextModule := context.(Module)

			m, err := contextModule.InstanceMethod(aliasNode.From.Name)
			if err != nil {
				returnErr = NewNameError(aliasNode.From.Name, contextModule.String(), contextModule.String(), vm.stack.String())
				return returnValue, returnErr
			}

			contextModule.AddInstanceMethod(NewNativeMethod(aliasNode.To.Name, vm, vm, func(self Value, block Block, args ...Value) (Value, error) {
				return m.Execute(self, block, args...)
			}))

		case ast.ModuleDecl:
			moduleNode := statement.(ast.ModuleDecl)
			theModule := NewModule(moduleNode.Name, vm, vm)
			vm.CurrentModules[moduleNode.Name] = theModule

			_, err := vm.executeWithContext(theModule, moduleNode.Body...)
			if err != nil {
				returnErr = err
			}

			returnValue = theModule

		case ast.ClassDecl:
			classNode := statement.(ast.ClassDecl)
			theClass := NewUserDefinedClass(classNode.Name, vm, vm)
			vm.CurrentClasses[classNode.FullName()] = theClass

			_, err := vm.executeWithContext(theClass, classNode.Body...)
			if err != nil {
				returnErr = err
			} else {
				returnValue = theClass
			}

		case ast.FuncDecl:
			funcNode := statement.(ast.FuncDecl)
			method := NewRubyMethod(
				funcNode.MethodName(),
				funcNode.MethodArgs(),
				funcNode.Body,
				vm,
				vm,
				func(self Value, method *RubyMethod) (Value, error) {
					vm.localVariableStack.unshift()
					defer vm.localVariableStack.shift()

					for _, arg := range method.Args() {
						vm.localVariableStack.store(arg.Name, arg.Value)
					}

					return vm.executeWithContext(self, method.Body()...)
				})
			returnValue = method

			if context == vm.ObjectSpace["main"] && funcNode.Target == nil {
				vm.CurrentModules["Kernel"].AddPrivateMethod(method)
			} else {
				switch funcNode.Target.(type) {
				case ast.Self:
					context.AddMethod(method)
				case nil:
					context.(Module).AddInstanceMethod(method)
				default:
					value, err := vm.executeWithContext(context, funcNode.Target)
					if err != nil {
						return nil, err
					}

					value.AddMethod(method)
				}
			}

		case ast.Nil:
			returnValue = vm.singletons["nil"]
		case ast.SimpleString:
			returnValue = NewString(statement.(ast.SimpleString).Value, vm, vm)
		case ast.InterpolatedString:
			returnValue = NewString(statement.(ast.InterpolatedString).Value, vm, vm)
		case ast.Boolean:
			if statement.(ast.Boolean).Value {
				returnValue = vm.singletons["true"]
			} else {
				returnValue = vm.singletons["false"]
			}
		case ast.GlobalVariable:
			returnValue = vm.CurrentGlobals[statement.(ast.GlobalVariable).Name]
		case ast.ConstantInt:
			returnValue = NewFixnum(statement.(ast.ConstantInt).Value, vm, vm)
		case ast.ConstantFloat:
			returnValue = NewFloat(statement.(ast.ConstantFloat).Value, vm)
		case ast.Symbol:
			name := statement.(ast.Symbol).Name
			maybe, ok := vm.CurrentSymbols[name]
			if !ok {
				returnValue = NewSymbol(name, vm)
				vm.CurrentSymbols[name] = returnValue
			} else {
				returnValue = maybe
			}
		case ast.BareReference:
			name := statement.(ast.BareReference).Name
			maybe, err := vm.localVariableStack.retrieve(name)
			if err == nil {
				returnValue = maybe
			} else {
				maybe, ok := vm.ObjectSpace[name]
				if ok {
					returnValue = maybe
				} else {
					maybe, ok := vm.CurrentClasses[name]
					if ok {
						returnValue = maybe
					} else {
						maybe, ok := vm.CurrentModules[name]
						if ok {
							returnValue = maybe
						} else {
							returnValue = nil
							returnErr = NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
						}
					}
				}
			}
		case ast.CallExpression:
			var method Method
			callExpr := statement.(ast.CallExpression)

			var (
				target           Value
				usePrivateMethod bool // FIXME: this should be unnecessary now
			)

			if callExpr.Target != nil {
				target, returnErr = vm.executeWithContext(context, callExpr.Target)
				if returnErr != nil {
					return nil, returnErr
				}
			} else {
				usePrivateMethod = true
				target = context
			}

			if target == nil {
				nilValue := vm.singletons["nil"]
				return nil, NewNoMethodError(callExpr.Func.Name, nilValue.String(), nilValue.Class().String(), vm.stack.String())
			}

			method, err := target.Method(callExpr.Func.Name)
			if err != nil && usePrivateMethod {
				method, err = target.PrivateMethod(callExpr.Func.Name)
			}

			if err != nil {
				return nil, err
			}

			args := []Value{}
			for _, astArgument := range callExpr.Args {
				arg, err := vm.executeWithContext(context, astArgument)
				if err != nil {
					return nil, err
				}

				args = append(args, arg)
			}

			vm.stack.Unshift(method.Name(), vm.currentFilename)
			defer vm.stack.Shift()

			var block Block
			if callExpr.OptionalBlock.Provided() {
				blockValue, err := vm.executeWithContext(context, callExpr.OptionalBlock)

				if err != nil {
					return nil, err
				}

				block = blockValue.(Block)
			}

			returnValue, returnErr = method.Execute(target, block, args...)
			if returnErr != nil {
				return returnValue, returnErr
			}

		case ast.Block:
			astBlock := statement.(ast.Block)
			block := NewBlock(context, astBlock.Args, astBlock.Body, vm)
			returnValue = block.(Value)

		case ast.Assignment:
			var err error
			assignment := statement.(ast.Assignment)
			returnValue, err = vm.executeWithContext(context, assignment.RHS)
			if err != nil {
				return nil, err
			}

			switch assignment.LHS.(type) {
			case ast.BareReference:
				ref := assignment.LHS.(ast.BareReference)
				vm.ObjectSpace[ref.Name] = returnValue
			case ast.GlobalVariable:
				globalVar := assignment.LHS.(ast.GlobalVariable)
				vm.CurrentGlobals[globalVar.Name] = returnValue
			case ast.InstanceVariable:
				iVar := assignment.LHS.(ast.InstanceVariable)
				context.SetInstanceVariable(iVar.Name, returnValue)
			default:
				panic(fmt.Sprintf("unimplemented assignment failure: %#v", assignment.LHS))
			}

		case ast.FileNameConstReference:
			returnValue = NewString(vm.currentFilename, vm, vm)
		case ast.Begin:
			begin := statement.(ast.Begin)
			_, err := vm.executeWithContext(context, begin.Body...)

			if err != nil {
				matchingRescue := false

				rubyErr, ok := err.(Value)
				if !ok {
					panic(context)
					return nil, err
				}

				for _, rescue := range begin.Rescue {
					if matchingRescue {
						break
					}

					r := rescue.(ast.Rescue)
					for _, exceptionClass := range r.Exception.Classes {
						if exceptionClass.Name == rubyErr.String() {
							_, err = vm.executeWithContext(context, r.Body...)
							if err == nil {
								matchingRescue = true
								break
							}
						}
					}
				}
			}

			if err != nil {
				returnErr = err
			}
		case ast.Array:
			arrayValue, _ := vm.CurrentClasses["Array"].New(vm, vm)
			array := arrayValue.(*Array)
			for _, node := range statement.(ast.Array).Nodes {
				value, err := vm.executeWithContext(context, node)
				if err != nil {
					return nil, err
				}

				array.Append(value)
			}

			returnValue = array

		case ast.Hash:
			hashValue, _ := vm.CurrentClasses["Hash"].New(vm, vm)
			hash := hashValue.(*Hash)
			for _, keyPair := range statement.(ast.Hash).Pairs {
				key, err := vm.executeWithContext(context, keyPair.Key)
				if err != nil {
					returnErr = err
					break
				}

				val, err := vm.executeWithContext(context, keyPair.Value)
				if err != nil {
					returnErr = err
					break
				}

				hash.Add(key, val)
			}

			returnValue = hash
		case ast.Ternary:
			ternary := statement.(ast.Ternary)
			value, err := vm.executeWithContext(context, ternary.Condition)
			if err != nil {
				returnErr = err
			} else {
				if value.IsTruthy() {
					returnValue, returnErr = vm.executeWithContext(context, ternary.True)
				} else {
					returnValue, returnErr = vm.executeWithContext(context, ternary.False)
				}

			}

		case ast.Class:
			class := statement.(ast.Class)
			className := class.FullName()
			value, ok := vm.CurrentClasses[className]
			if !ok {
				returnErr = NewNameError(className, context.String(), context.Class().String(), vm.stack.String())
			} else {
				returnValue = value
			}

		default:
			panic(fmt.Sprintf("handled unknown statement type: %T:\n\t\n => %#v\n", statement, statement))
		}
	}

	return returnValue, returnErr
}

// ClassProvider
func (vm *vm) ClassWithName(name string) Class {
	return vm.CurrentClasses[name]
}

// ArgEvaluator
func (vm *vm) EvaluateArgInContext(arg ast.Node, context Value) (Value, error) {
	return vm.executeWithContext(context, arg)
}

// BlockEvaluator
func (vm *vm) EvaluateBlockWithArgsInContext(
	context Value,
	args []BlockArg,
	statements []ast.Node) (Value, error) {
	vm.localVariableStack.unshift()
	defer vm.localVariableStack.shift()

	for _, arg := range args {
		vm.localVariableStack.store(arg.Name, arg.Value)
	}

	return vm.executeWithContext(context, statements...)
}

// SingletonProvider
func (vm *vm) SingletonWithName(name string) Value {
	return vm.singletons[name]
}

func (vm *vm) NewSingletonWithName(name string, value Value) {
	vm.singletons[name] = value
}

func (vm *vm) SymbolWithName(name string) Value {
	return vm.CurrentSymbols[name]
}

func (vm *vm) AddSymbol(val Value) {
	symbol := val.(*SymbolValue)
	vm.CurrentSymbols[symbol.Name()] = symbol
}
