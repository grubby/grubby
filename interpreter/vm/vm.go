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

	ObjectSpace    map[string]Value
	CurrentGlobals map[string]Value
	CurrentSymbols map[string]Value
	CurrentClasses map[string]Class
	CurrentModules map[string]Module
	singletons     map[string]Value

	stack              *CallStack
	localVariableStack *LocalVariableStack

	inEigenclassBlock bool

	currentModuleName string // FIXME: I bet this could be determined from context

	exitCallbacks []Block

	required_files map[string]bool
}

type VM interface {
	Run(string) (Value, error)
	Exit()

	Get(string) (Value, error)
	MustGet(string) Value

	GetClass(string) (Class, error)
	MustGetClass(string) Class

	GetModule(string) (Module, error)
	MustGetModule(string) Module

	Set(string, Value)

	Symbols() map[string]Value
	Globals() map[string]Value
	Classes() map[string]Class
	Modules() map[string]Module

	Provider
	ArgEvaluator
	ClassProvider
	SingletonProvider
	StackProvider
}

func NewVM(rubyHome, name string) VM {
	vm := &vm{
		currentFilename:    name,
		stack:              NewCallStack(),
		CurrentGlobals:     make(map[string]Value),
		ObjectSpace:        make(map[string]Value),
		CurrentSymbols:     make(map[string]Value),
		CurrentModules:     make(map[string]Module),
		localVariableStack: NewLocalVariableStack(),
		singletons:         make(map[string]Value),
		required_files:     make(map[string]bool),
	}
	vm.registerBuiltinClassesAndModules()

	loadPath, _ := vm.CurrentClasses["Array"].New(vm)
	loadPath.(*Array).Append(NewString(filepath.Join(rubyHome, "lib"), vm))

	vm.CurrentGlobals["LOAD_PATH"] = loadPath
	vm.CurrentGlobals[":"] = loadPath

	argvArray, err := vm.CurrentClasses["Array"].New(vm)
	if err != nil {
		panic(err)
	}
	vm.CurrentClasses["Object"].SetConstant("ARGV", argvArray)
	vm.CurrentClasses["Object"].SetConstant("RUBY_NAME", NewString("grubby", vm))
	vm.CurrentClasses["Object"].SetConstant("ENV", NewENVConstant(vm))

	// this is a temporary hack to gain progress on running rubyspec
	// as an alternative, we could implement String#=~
	vm.CurrentClasses["Object"].SetConstant("RUBY_EXE", NewString("bin/ruby", vm))

	main, _ := vm.CurrentClasses["Object"].New(vm)
	main.AddMethod(NewNativeMethod("to_s", vm, func(self Value, block Block, args ...Value) (Value, error) {
		return NewString("main", vm), nil
	}))
	vm.ObjectSpace["main"] = main

	return vm
}

func (vm *vm) registerBuiltinClassesAndModules() {
	vm.CurrentClasses = map[string]Class{}
	vm.CurrentModules = map[string]Module{}

	basicObjectClass := NewBasicObjectClass(vm)
	vm.CurrentClasses["BasicObject"] = basicObjectClass

	objectClass := NewGlobalObjectClass(vm)
	vm.CurrentClasses["Object"] = objectClass

	classClass := NewClassClass(vm)
	vm.CurrentClasses["Class"] = classClass

	moduleClass := NewModuleClass(vm, vm)
	vm.CurrentClasses["Module"] = moduleClass
	vm.CurrentModules["Comparable"] = NewComparableModule(vm)
	vm.CurrentModules["Kernel"] = NewGlobalKernelModule(vm)
	vm.CurrentModules["Process"] = NewProcessModule(vm)

	// FIXME: this should be private, but method resolution fails
	vm.CurrentModules["Kernel"].AddMethod(NewNativeMethod("require", vm, func(self Value, block Block, args ...Value) (Value, error) {
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

			absolutePath, err := filepath.Abs(fullPath)
			if err != nil {
				return nil, err
			}

			_, ok := vm.required_files[absolutePath]
			if ok {
				return vm.singletons["false"], nil
			}
			vm.required_files[absolutePath] = true

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
	vm.CurrentModules["Kernel"].AddMethod(NewNativeMethod("at_exit", vm, func(self Value, block Block, args ...Value) (Value, error) {
		if block != nil {
			vm.exitCallbacks = append(vm.exitCallbacks, block)
		}

		return nil, nil
	}))
	vm.CurrentModules["Kernel"].AddMethod(NewNativeMethod("object_id", vm, func(self Value, block Block, args ...Value) (Value, error) {
		return NewFixnum(1, vm), nil
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

	vm.CurrentClasses["NilClass"] = NewNilClass(vm)
	vm.CurrentClasses["TrueClass"] = NewTrueClass(vm)
	vm.CurrentClasses["FalseClass"] = NewFalseClass(vm)

	vm.singletons["nil"], _ = vm.CurrentClasses["NilClass"].New(vm)
	vm.singletons["true"], _ = vm.CurrentClasses["TrueClass"].New(vm)
	vm.singletons["false"], _ = vm.CurrentClasses["FalseClass"].New(vm)

	vm.CurrentClasses["IO"] = NewIOClass(vm)
	vm.CurrentClasses["Array"] = NewArrayClass(vm)
	vm.CurrentClasses["Hash"] = NewHashClass(vm)
	vm.CurrentClasses["String"] = NewStringClass(vm)
	vm.CurrentClasses["Numeric"] = NewNumericClass(vm)
	vm.CurrentClasses["Integer"] = NewIntegerClass(vm)
	vm.CurrentClasses["Fixnum"] = NewFixnumClass(vm)
	vm.CurrentClasses["Float"] = NewFloatClass(vm)
	vm.CurrentClasses["Symbol"] = NewSymbolClass(vm)
	vm.CurrentClasses["Proc"] = NewProcClass(vm)
	vm.CurrentClasses["Regexp"] = NewRegexpClass(vm)
	vm.CurrentClasses["File"] = NewFileClass(vm)
	vm.CurrentClasses["Dir"] = NewDirClass(vm)
	vm.CurrentClasses["Exception"] = NewExceptionClass(vm)
	vm.CurrentClasses["StandardError"] = NewStandardErrorClass(vm)
	vm.CurrentClasses["ArgumentError"] = NewArgumentErrorClass(vm)
	vm.CurrentClasses["Encoding"] = NewEncodingClass(vm)
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

func (vm *vm) GetModule(name string) (Module, error) {
	for keyName, module := range vm.CurrentModules {
		if keyName == name {
			return module, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Module '%s' not found", name))
}

func (vm *vm) MustGetModule(name string) Module {
	module, err := vm.GetModule(name)
	if err != nil {
		panic(fmt.Sprintf("module '%s' requested, but does not exist", name))
	}

	return module
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
	parser.Reset()

	lexer := parser.NewLexer(input)
	result := parser.RubyParse(lexer)
	if result != 0 {
		return nil, NewParseError(vm.currentFilename)
	}

	main := vm.ObjectSpace["main"]
	vm.stack.Unshift("main", vm.currentFilename, 0)
	defer vm.stack.Shift()

	vm.localVariableStack.Unshift()
	defer vm.localVariableStack.Shift()
	return vm.executeWithContext(main, parser.Statements...)
}

func (vm *vm) Exit() {
	for _, block := range vm.exitCallbacks {
		block.Call()
	}
}

func (vm *vm) executeWithContext(context Value, statements ...ast.Node) (Value, error) {
	var (
		returnValue Value
		returnErr   error
	)
	for _, statement := range statements {
		if returnErr != nil {
			break
		}

		switch statement.(type) {
		case ast.Self:
			returnValue = context
		case ast.Return:
			returnNode := statement.(ast.Return)
			return vm.executeWithContext(context, returnNode.Value)
		case ast.IfBlock:
			returnValue, returnErr = interpretIfStatementInContext(vm, statement.(ast.IfBlock), context)
		case ast.Alias:
			returnValue, returnErr = interpretAliasInContext(vm, statement.(ast.Alias), context)
		case ast.ModuleDecl:
			returnValue, returnErr = interpretModuleDeclarationInContext(vm, statement.(ast.ModuleDecl), context)
		case ast.ClassDecl:
			returnValue, returnErr = interpretClassDeclarationInContext(vm, statement.(ast.ClassDecl), context)
		case ast.FuncDecl:
			returnValue, returnErr = interpretMethodDeclarationInContext(vm, statement.(ast.FuncDecl), context)
		case ast.Nil:
			returnValue = vm.singletons["nil"]
		case ast.SimpleString:
			returnValue = NewString(statement.(ast.SimpleString).Value, vm)
		case ast.InterpolatedString:
			returnValue, returnErr = interpretDoubleQuotedStringInContext(vm, statement.(ast.InterpolatedString), context)
		case ast.Boolean:
			if statement.(ast.Boolean).Value {
				returnValue = vm.singletons["true"]
			} else {
				returnValue = vm.singletons["false"]
			}
		case ast.GlobalVariable:
			var ok bool
			returnValue, ok = vm.CurrentGlobals[statement.(ast.GlobalVariable).Name]
			if !ok {
				returnValue = vm.singletons["nil"]
			}
		case ast.ConstantInt:
			returnValue = NewFixnum(statement.(ast.ConstantInt).Value, vm)
		case ast.ConstantFloat:
			returnValue = NewFloat(statement.(ast.ConstantFloat).Value, vm)
		case ast.Symbol:
			returnValue = interpretSymbol(vm, statement.(ast.Symbol))
		case ast.BareReference:
			returnValue, returnErr = interpretBareReferenceInContext(vm, statement.(ast.BareReference), context)
		case ast.CallExpression:
			returnValue, returnErr = interpretCallExpressionInContext(vm, statement.(ast.CallExpression), context)
		case ast.Block:
			astBlock := statement.(ast.Block)
			block := NewBlock(context, astBlock.Args, astBlock.Body, vm)
			returnValue = block.(Value)

		case ast.Assignment:
			returnValue, returnErr = interpretAssignmentInContext(vm, statement.(ast.Assignment), context)
		case ast.FileNameConstReference:
			returnValue = NewString(vm.currentFilename, vm)
		case ast.LineNumberConstReference:
			lineNumber := int64(statement.(ast.LineNumberConstReference).LineNumber())
			returnValue = NewFixnum(lineNumber, vm)
		case ast.Begin:
			returnValue, returnErr = interpretBeginInContext(vm, statement.(ast.Begin), context)
		case ast.Array:
			returnValue, returnErr = interpretArrayInContext(vm, statement.(ast.Array), context)
		case ast.Hash:
			returnValue, returnErr = interpretHashInContext(vm, statement.(ast.Hash), context)
		case ast.Ternary:
			returnValue, returnErr = interpretTernaryInContext(vm, statement.(ast.Ternary), context)
		case ast.Class:
			returnValue, returnErr = interpretClassInContext(vm, statement.(ast.Class), context)
		case ast.Eigenclass:
			returnValue, returnErr = interpretEigenclassInContext(vm, statement.(ast.Eigenclass), context)
		case ast.Constant:
			returnValue, returnErr = interpretConstantInContext(vm, statement.(ast.Constant), context)
		case ast.Negation:
			returnValue, returnErr = interpretNegationInContext(vm, statement.(ast.Negation), context)
		case ast.Regex:
			returnValue, returnErr = interpretRegexpInContext(vm, statement.(ast.Regex), context)
		case ast.WeakLogicalAnd:
			returnValue, returnErr = interpretWeakBooleanAnd(vm, statement.(ast.WeakLogicalAnd), context)
		case ast.WeakLogicalOr:
			returnValue, returnErr = interpretWeakBooleanOr(vm, statement.(ast.WeakLogicalOr), context)

		case ast.Defined:
			returnValue, returnErr = interpretDefinedKeyword(vm, statement.(ast.Defined), context)
		case ast.InstanceVariable:
			returnValue, returnErr = interpretInstanceVariableInContext(vm, statement.(ast.InstanceVariable), context)
		case ast.SwitchStatement:
			returnValue, returnErr = interpretSwitchStatement(vm, statement.(ast.SwitchStatement), context)
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
	vm.localVariableStack.UnshiftCopyingCurrentFrame()
	defer vm.localVariableStack.Shift()

	for _, arg := range args {
		vm.localVariableStack.Store(arg.Name, arg.Value)
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

// Evaluator
func (vm *vm) EvaluateStringInContext(input string, context Value) (Value, error) {
	parser.Reset()

	lexer := parser.NewLexer(input)
	result := parser.RubyParse(lexer)
	if result != 0 {
		return nil, NewParseError(vm.currentFilename)
	}

	return vm.executeWithContext(context, parser.Statements...)
}

func (vm *vm) EvaluateStringInContextAndNewStack(input string, context Value) (Value, error) {
	parser.Reset()

	lexer := parser.NewLexer(input)
	result := parser.RubyParse(lexer)
	if result != 0 {
		return nil, NewParseError(vm.currentFilename)
	}

	vm.localVariableStack.Unshift()
	defer vm.localVariableStack.Shift()
	return vm.executeWithContext(context, parser.Statements...)
}

// StackProvider
func (vm *vm) CurrentStack() string {
	return vm.stack.String()
}

// MethodProvider
func (vm *vm) AddMethod(name string, context Value, body func(self Value, block Block, args ...Value) (Value, error)) {
	if vm.inEigenclassBlock {
		context.AddMethod(NewNativeMethod(name, vm, body))
	} else {
		context.(Module).AddInstanceMethod(NewNativeMethod(name, vm, body))
	}
}

// General Purpose Provider
func (vm *vm) ArgEvaluator() ArgEvaluator {
	return vm
}
func (vm *vm) ClassProvider() ClassProvider {
	return vm
}
func (vm *vm) SingletonProvider() SingletonProvider {
	return vm
}
func (vm *vm) StackProvider() StackProvider {
	return vm
}
func (vm *vm) MethodProvider() MethodProvider {
	return vm
}
