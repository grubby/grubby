package vm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/interpreter/vm/builtins"
	"github.com/grubby/grubby/parser"
)

type vm struct {
	currentFilename string

	stack          *CallStack
	ObjectSpace    map[string]builtins.Value
	CurrentGlobals map[string]builtins.Value
	CurrentSymbols map[string]builtins.Value
	CurrentClasses map[string]builtins.Class
	CurrentModules map[string]builtins.Module

	localVariableStack *localVariableStack
}

type VM interface {
	Run(string) (builtins.Value, error)
	Get(string) (builtins.Value, error)
	MustGet(string) builtins.Value

	GetClass(string) (builtins.Class, error)
	MustGetClass(string) builtins.Class

	Set(string, builtins.Value)

	Symbols() map[string]builtins.Value
	Globals() map[string]builtins.Value
	Classes() map[string]builtins.Class
	Modules() map[string]builtins.Module

	builtins.ClassProvider
}

func NewVM(rubyHome, name string) VM {
	vm := &vm{
		currentFilename:    name,
		stack:              NewCallStack(),
		CurrentGlobals:     make(map[string]builtins.Value),
		ObjectSpace:        make(map[string]builtins.Value),
		CurrentSymbols:     make(map[string]builtins.Value),
		CurrentModules:     make(map[string]builtins.Module),
		localVariableStack: newLocalVariableStack(),
	}
	vm.registerBuiltinClassesAndModules()

	loadPath := vm.CurrentClasses["Array"].New(vm)
	loadPath.(*builtins.Array).Append(builtins.NewString(filepath.Join(rubyHome, "lib"), vm))

	vm.CurrentGlobals["LOAD_PATH"] = loadPath
	vm.CurrentGlobals[":"] = loadPath
	vm.ObjectSpace["ARGV"] = vm.CurrentClasses["Array"].New(vm)

	main := vm.CurrentClasses["Object"].New(vm)
	main.AddMethod(builtins.NewNativeMethod("to_s", vm, func(self builtins.Value, args ...builtins.Value) (builtins.Value, error) {
		return builtins.NewString("main", vm), nil
	}))
	main.AddMethod(builtins.NewNativeMethod("require", vm, func(self builtins.Value, args ...builtins.Value) (builtins.Value, error) {
		fileName := args[0].(*builtins.StringValue).String()
		if fileName == "rubygems" {
			// don't "require 'rubygems'"
			return vm.CurrentClasses["False"].New(vm), nil
		}

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
				_, rubyErr := vm.Run(string(contents))
				return vm.CurrentClasses["True"].New(vm), rubyErr
			}
		}

		errorMessage := fmt.Sprintf("LoadError: cannot load such file -- %s", fileName)
		return nil, builtins.NewLoadError(errorMessage, vm.stack.String())
	}))
	vm.ObjectSpace["main"] = main

	return vm
}

func (vm *vm) registerBuiltinClassesAndModules() {
	vm.CurrentClasses = map[string]builtins.Class{}
	vm.CurrentModules = map[string]builtins.Module{}

	basicObjectClass := builtins.NewBasicObjectClass(vm)
	vm.CurrentClasses["BasicObject"] = basicObjectClass

	objectClass := builtins.NewGlobalObjectClass(vm)
	vm.CurrentClasses["Object"] = objectClass

	// FIXME: need to set Class' superclass to module at this point
	// and need to set "class" as the class on object and basic object
	classClass := builtins.NewClassClass(vm)
	vm.CurrentClasses["Class"] = classClass

	moduleClass := builtins.NewModuleClass(vm)
	vm.CurrentClasses["Module"] = moduleClass
	vm.CurrentModules["Comparable"] = builtins.NewComparableModule(vm)
	vm.CurrentModules["Kernel"] = builtins.NewGlobalKernelModule(vm)
	vm.CurrentModules["Process"] = builtins.NewProcessModule(vm)

	/* BEGIN RUNTIME TRICKERY
	There's a cycle in ruby's builtin object graph
	There are classes that refer to each other (Module, Class)
	As well as including the Kernel module
	*/
	objectClass.Include(vm.CurrentModules["Kernel"])
	moduleClass.Include(vm.CurrentModules["Kernel"])
	classClass.(*builtins.ClassValue).SetSuperClass()
	objectClass.(*builtins.ObjectClass).SetSuperClass()
	basicObjectClass.(*builtins.BasicObjectClass).SetSuperClass()
	// END RUNTIME TRICKERY

	vm.CurrentClasses["IO"] = builtins.NewIOClass(vm)
	vm.CurrentClasses["Array"] = builtins.NewArrayClass(vm)
	vm.CurrentClasses["Hash"] = builtins.NewHashClass(vm)
	vm.CurrentClasses["True"] = builtins.NewTrueClass(vm)
	vm.CurrentClasses["File"] = builtins.NewFileClass(vm)
	vm.CurrentClasses["False"] = builtins.NewFalseClass(vm)
	vm.CurrentClasses["Nil"] = builtins.NewNilClass(vm)
	vm.CurrentClasses["String"] = builtins.NewStringClass(vm)
}

func (vm *vm) MustGet(key string) builtins.Value {
	val, err := vm.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func (vm *vm) Get(key string) (builtins.Value, error) {
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

func (vm *vm) GetClass(name string) (builtins.Class, error) {
	for _, class := range vm.CurrentClasses {
		if class.Name() == name {
			return class, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Class '%s' not found", name))
}

func (vm *vm) MustGetClass(name string) builtins.Class {
	for _, class := range vm.CurrentClasses {
		if class.Name() == name {
			return class
		}
	}

	panic(fmt.Sprintf("class '%s' requested, but does not exist", name))
}

func (vm *vm) Set(key string, value builtins.Value) {
	vm.ObjectSpace[key] = value
}

func (vm *vm) Symbols() map[string]builtins.Value {
	return vm.CurrentSymbols
}

func (vm *vm) Globals() map[string]builtins.Value {
	return vm.CurrentGlobals
}

func (vm *vm) Classes() map[string]builtins.Class {
	return vm.CurrentClasses
}

func (vm *vm) Modules() map[string]builtins.Module {
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

func (vm *vm) Run(input string) (builtins.Value, error) {
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

func (vm *vm) executeWithContext(context builtins.Value, statements ...ast.Node) (builtins.Value, error) {
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
				returnValue, returnErr = vm.executeWithContext(context, ifBlock.Body...)
			} else {
				returnValue, returnErr = vm.executeWithContext(context, ifBlock.Else...)
			}
		case ast.Alias:
			// FIXME: assumes that the context will be a module, but could also be a class
			aliasNode := statement.(ast.Alias)
			contextModule := context.(*builtins.RubyModule)

			m, err := contextModule.InstanceMethod(aliasNode.From.Name)
			if err != nil {
				returnErr = builtins.NewNameError(aliasNode.From.Name, contextModule.String(), contextModule.String(), vm.stack.String())
				return returnValue, returnErr
			}

			contextModule.AddInstanceMethod(builtins.NewNativeMethod(aliasNode.To.Name, vm, func(self builtins.Value, args ...builtins.Value) (builtins.Value, error) {
				return m.Execute(self, args...)
			}))

		case ast.ModuleDecl:
			moduleNode := statement.(ast.ModuleDecl)
			theModule := builtins.NewModule(moduleNode.Name, vm)
			vm.CurrentModules[moduleNode.Name] = theModule

			_, err := vm.executeWithContext(theModule, moduleNode.Body...)
			if err != nil {
				returnErr = err
			}

			returnValue = theModule

		case ast.ClassDecl:
			classNode := statement.(ast.ClassDecl)
			theClass := builtins.NewUserDefinedClass(classNode.Name, vm)
			vm.CurrentClasses[classNode.Name] = theClass

			_, err := vm.executeWithContext(theClass, classNode.Body...)
			if err != nil {
				returnErr = err
			}

		case ast.FuncDecl:
			funcNode := statement.(ast.FuncDecl)
			method := builtins.NewRubyMethod(funcNode.MethodName(), funcNode.MethodArgs(), funcNode.Body, vm, func(self builtins.Value, method *builtins.RubyMethod) (builtins.Value, error) {
				vm.localVariableStack.unshift()
				defer vm.localVariableStack.shift()

				for _, arg := range method.Args() {
					vm.localVariableStack.store(arg.Name, arg.Value)
				}

				return vm.executeWithContext(context, method.Body()...)
			})
			returnValue = method

			if context == vm.ObjectSpace["main"] {
				vm.CurrentModules["Kernel"].AddPrivateMethod(method)
			} else {
				switch context.(type) {
				case builtins.Class:
					context.(builtins.Class).AddInstanceMethod(method)
				case builtins.Module:
					ref, ok := funcNode.Target.(ast.BareReference)
					if ok && ref.Name == "self" {
						context.AddMethod(method)
					} else {
						context.(builtins.Module).AddInstanceMethod(method)
					}
				default:
					panic(fmt.Sprintf("unknown type of context: %#T", context))
				}
			}

		case ast.SimpleString:
			returnValue = builtins.NewString(statement.(ast.SimpleString).Value, vm)
		case ast.InterpolatedString:
			returnValue = builtins.NewString(statement.(ast.InterpolatedString).Value, vm)
		case ast.Boolean:
			if statement.(ast.Boolean).Value {
				returnValue = vm.CurrentClasses["True"].New(vm)
			} else {
				returnValue = vm.CurrentClasses["False"].New(vm)
			}
		case ast.GlobalVariable:
			returnValue = vm.CurrentGlobals[statement.(ast.GlobalVariable).Name]
		case ast.ConstantInt:
			returnValue = builtins.NewInt(statement.(ast.ConstantInt).Value)
		case ast.ConstantFloat:
			returnValue = builtins.NewFloat(statement.(ast.ConstantFloat).Value)
		case ast.Symbol:
			name := statement.(ast.Symbol).Name
			maybe, ok := vm.CurrentSymbols[name]
			if !ok {
				returnValue = builtins.NewSymbol(name)
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
							returnErr = builtins.NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
						}
					}
				}
			}
		case ast.CallExpression:
			var method builtins.Method
			callExpr := statement.(ast.CallExpression)

			var (
				target           builtins.Value
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
				nilValue := vm.CurrentClasses["Nil"].New(vm)
				return nil, builtins.NewNoMethodError(callExpr.Func.Name, nilValue.String(), nilValue.Class().String(), vm.stack.String())
			}

			method, err := target.Method(callExpr.Func.Name)
			if err != nil && usePrivateMethod {
				method, err = target.PrivateMethod(callExpr.Func.Name)
			}

			if err != nil {
				return nil, err
			}

			args := []builtins.Value{}
			for _, astArgument := range callExpr.Args {
				arg, err := vm.executeWithContext(context, astArgument)
				if err != nil {
					return nil, err
				}

				args = append(args, arg)
			}

			vm.stack.Unshift(method.Name(), vm.currentFilename)
			defer vm.stack.Shift()

			returnValue, returnErr = method.Execute(target, args...)
			if returnErr != nil {
				return returnValue, returnErr
			}

		case ast.Assignment:
			assignment := statement.(ast.Assignment)
			returnValue, err := vm.executeWithContext(context, assignment.RHS)
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
			default:
				panic(fmt.Sprintf("unimplemented assignment failure: %#v", assignment.LHS))
			}

		case ast.FileNameConstReference:
			returnValue = builtins.NewString(vm.currentFilename, vm)
		case ast.Begin:
			begin := statement.(ast.Begin)
			_, err := vm.executeWithContext(context, begin.Body...)

			if err != nil {
				matchingRescue := false
				rubyErr := err.(builtins.Value)
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
			array := vm.CurrentClasses["Array"].New(vm).(*builtins.Array)
			for _, node := range statement.(ast.Array).Nodes {
				value, err := vm.executeWithContext(context, node)
				if err != nil {
					return nil, err
				}

				array.Append(value)
			}

			returnValue = array

		case ast.Hash:
			hash := vm.CurrentClasses["Hash"].New(vm).(*builtins.Hash)
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
		default:
			panic(fmt.Sprintf("handled unknown statement type: %T:\n\t\n => %#v\n", statement, statement))
		}
	}

	return returnValue, returnErr
}

func (vm *vm) ClassWithName(name string) builtins.Class {
	return vm.CurrentClasses[name]
}
