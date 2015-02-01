package vm_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VM", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Describe("the global Object", func() {
		It("exists", func() {
			Expect(vm.MustGet("Object")).ToNot(BeNil())
		})

		It("reports its class as being 'Class'", func() {
			o := vm.MustGet("Object")
			Expect(o.Class().String()).To(Equal("Class"))
		})
	})

	Describe("creating a simple function", func() {
		BeforeEach(func() {
			_, err := vm.Run(`
def foo
end`)

			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a private method on Kernel", func() {
			kernel := vm.Modules()["Kernel"]
			method, err := kernel.PrivateMethod("foo")

			Expect(err).ToNot(HaveOccurred())
			Expect(method.Name()).To(Equal("foo"))
		})

		// FIXME: to fix this, methods should know about
		// * superclass methods (ideally the entire chain)
		// * included modules
		// * extended modules
		// (((   and Object needs to "include" Kernel   )))
		PIt("is also available in Object", func() {
			object, err := vm.Get("Object")
			Expect(err).ToNot(HaveOccurred())
			Expect(object).To(HaveMethod("foo"))
		})
	})

	Describe("strings", func() {
		It("returns a ruby String object", func() {
			val, err := vm.Run("'nonrestricted-consonantize'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewString("", vm, vm)))
			Expect(val.String()).To(Equal(`"nonrestricted-consonantize"`))
		})

		It("has a + method", func() {
			val, err := vm.Run("'foo' + 'bar'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal(builtins.NewString("foobar", vm, vm).String()))
		})
	})

	Describe("numbers", func() {
		It("returns a ruby Fixnum for a number literal", func() {
			val, err := vm.Run("5")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewFixnum(0, vm, vm)))
			Expect(val.String()).To(Equal("5"))
			Expect(val.Class()).To(Equal(vm.MustGetClass("Fixnum")))
		})

		It("treats numbers as singletons", func() {
			val, err := vm.Run("5 == 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(Equal(vm.SingletonWithName("true")))
		})
	})

	Describe("interpreting a float", func() {
		It("returns a ruby Float object", func() {
			val, err := vm.Run("5.123")

			Expect(err).ToNot(HaveOccurred())

			Expect(val).To(BeAssignableToTypeOf(builtins.NewFloat(0.0, vm)))
			Expect(val.Class()).To(Equal(vm.MustGetClass("Float")))

			asFloat, ok := val.(*builtins.FloatValue)
			Expect(ok).To(BeTrue())
			Expect(asFloat.ValueAsFloat()).To(Equal(5.123))
		})
	})

	Describe("interpreting a symbol", func() {
		var (
			val builtins.Value
			err error
		)

		BeforeEach(func() {
			val, err = vm.Run(":foo")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an instance of the Symbol class", func() {
			Expect(val.Class()).To(Equal(vm.MustGetClass("Symbol")))
		})

		It("returns a ruby Symbol object", func() {
			Expect(val).To(Equal(vm.Symbols()["foo"]))
		})

		It("registers the symbol globally", func() {
			Expect(val).To(Equal(vm.Symbols()["foo"]))
		})

		It("records the symbol only once", func() {
			sameSymbol, err := vm.Run(":foo")
			Expect(err).ToNot(HaveOccurred())

			firstPointer := reflect.ValueOf(val).Pointer()
			secondPointer := reflect.ValueOf(sameSymbol).Pointer()
			Expect(secondPointer).To(Equal(firstPointer))
		})
	})

	Describe("nil", func() {
		var (
			val builtins.Value
			err error
		)

		BeforeEach(func() {
			val, err = vm.Run("nil")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an instance of the NilClass class", func() {
			Expect(val.Class()).To(Equal(vm.MustGetClass("NilClass")))
		})

		It("treats nil as a singleton", func() {
			nilValue, err := vm.Run("nil")
			Expect(err).ToNot(HaveOccurred())

			firstPointer := reflect.ValueOf(val).Pointer()
			secondPointer := reflect.ValueOf(nilValue).Pointer()
			Expect(firstPointer).To(Equal(secondPointer))
		})
	})

	Describe("a reference to a variable", func() {
		Context("when it already is defined", func() {
			BeforeEach(func() {
				vm.Set("foo", builtins.NewString("superinquisitive-edacity", vm, vm))
			})

			It("returns the variable referenced", func() {
				value, err := vm.Run("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal(`"superinquisitive-edacity"`))
			})
		})

		Context("when it has not been defined", func() {
			It("returns an error", func() {
				_, err := vm.Run("grigri_incircumscription")

				Expect(err).To(HaveOccurred())
			})
		})
	})

	PDescribe("the standard lib", func() {
		It("is available to require", func() {
			_, err := vm.Run("require 'fileutils'")
			Expect(err).ToNot(BeAssignableToTypeOf(builtins.NewLoadError("", "")))
		})
	})

	Describe("Kernel#puts", func() {
		It("takes the given arguments and prints them to stdout", func() {
			output := SwapStdout(func() {
				_, err := vm.Run("puts 'conga-oestradiol'")
				Expect(err).ToNot(HaveOccurred())
			})

			Expect(output).To(ContainSubstring("conga-oestradiol"))
		})
	})

	Describe("Kernel#require", func() {
		It("searches for a file with the given name", func() {
			_, err := vm.Run("require 'something'")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(builtins.NewLoadError("", "")))
		})

		Context("with a load path and a file to require", func() {
			BeforeEach(func() {
				SetupLoadPathWithAFileToRequire(vm)
			})

			It("requires the file", func() {
				_, err := vm.Run("require 'foo'")
				Expect(err).ToNot(HaveOccurred())

				kernel := vm.Modules()["Kernel"]
				method, err := kernel.PrivateMethod("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(method.Name()).To(Equal("foo"))
			})
		})
	})

	Describe("the load path", func() {
		It("is represented by $LOAD_PATH and $:", func() {
			path := vm.MustGet("LOAD_PATH")
			str := builtins.NewString("foo", vm, vm)
			path.(*builtins.Array).Append(str)

			Expect(vm.MustGet(":").(*builtins.Array).Members()).To(ContainElement(str))
		})
	})

	Describe("Hash class", func() {
		It("it can be constructed from a hash literal", func() {
			hash, err := vm.Run("{:key => :value}")
			Expect(err).ToNot(HaveOccurred())

			keysMethod, err := hash.Method("keys")
			Expect(err).ToNot(HaveOccurred())

			keys, err := keysMethod.Execute(hash, nil)
			Expect(err).ToNot(HaveOccurred())

			keyArray, ok := keys.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(keyArray.Members()).To(ContainElement(vm.Symbols()["key"]))

			valuesMethod, err := hash.Method("values")
			Expect(err).ToNot(HaveOccurred())

			values, err := valuesMethod.Execute(hash, nil)
			Expect(err).ToNot(HaveOccurred())

			valueArray, ok := values.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(valueArray.Members()).To(ContainElement(vm.Symbols()["value"]))

			method, err := hash.Method("[]=")
			Expect(err).ToNot(HaveOccurred())

			fooSymbol := builtins.NewSymbol("foo", vm)
			barSymbol := builtins.NewSymbol("bar", vm)

			_, err = method.Execute(hash, nil, fooSymbol, barSymbol)
			Expect(err).ToNot(HaveOccurred())

			keys, err = keysMethod.Execute(hash, nil)
			Expect(err).ToNot(HaveOccurred())

			keyArray, ok = keys.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(keyArray.Members()).To(ContainElement(fooSymbol))

			values, err = valuesMethod.Execute(hash, nil)
			Expect(err).ToNot(HaveOccurred())

			valueArray, ok = values.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(valueArray.Members()).To(ContainElement(barSymbol))
		})
	})

	Describe("File class", func() {
		It("has a reasonable .expand_path method", func() {
			fileClass := vm.ClassWithName("File")
			Expect(fileClass).ToNot(BeNil())

			method, err := fileClass.Method("expand_path")
			Expect(err).ToNot(HaveOccurred())

			result, err := method.Execute(fileClass, nil, builtins.NewString("~/foobar", vm, vm))
			Expect(err).ToNot(HaveOccurred())

			expectedPath := fmt.Sprintf(`"%s/%s"`, os.Getenv("HOME"), "foobar")
			Expect(result.String()).To(Equal(expectedPath))
		})
	})

	Describe("assignment to a variable", func() {
		It("stores the value assigned", func() {
			_, err := vm.Run("foo = 'albitite-compotor'")
			Expect(err).ToNot(HaveOccurred())

			value, err := vm.Get("foo")
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*builtins.StringValue).RawString()).To(Equal("albitite-compotor"))
		})
	})

	Describe("special global variables", func() {
		Describe("__FILE__", func() {
			It("inherits the name given to the vm initially", func() {
				value, err := vm.Run("__FILE__")

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal(`"fake-irb-under-test"`))
			})

			It("uses the relative path to the file if used in a require'd file", func() {
				SetupFileWithGlobalFilenameConst(vm)
				_, err := vm.Run("require 'foo'")
				Expect(err).ToNot(HaveOccurred())

				value, err := vm.Get("foo")
				Expect(err).ToNot(HaveOccurred())

				// should this actually be the absolute path to foo.rb?
				Expect(value.String()).To(ContainSubstring("foo.rb"))
			})
		})
	})

	Describe("ARGV", func() {
		It("has a shift method", func() {
			value, err := vm.Run("ARGV.shift")
			Expect(err).ToNot(HaveOccurred())

			nilInstance := vm.SingletonWithName("nil")
			Expect(value).To(Equal(nilInstance))
		})
	})

	Describe("begin; rescue; end", func() {
		It("can be used to prevent exceptions from bubbling up", func() {
			_, err := vm.Run(`
foo = false
bar = false
begin
  require 'some/nonsense'
rescue LoadError
  foo = true
end

bar = true
`)

			Expect(err).ToNot(HaveOccurred())

			trueValue := vm.SingletonWithName("true")

			Expect(err).ToNot(HaveOccurred())
			Expect(vm.MustGet("foo")).To(Equal(trueValue))
			Expect(vm.MustGet("bar")).To(Equal(trueValue))
		})
	})

	Describe("calling a method that does not exist", func() {
		It("raises a NoMethodError", func() {
			_, err := vm.Run("'hello'.world()")
			Expect(err).To(BeAssignableToTypeOf(builtins.NewNoMethodError("", "", "", "")))
			Expect(err.Error()).To(ContainSubstring("undefined method 'world' for \"hello\":String"))
		})
	})

	Describe("stack traces", func() {
		return
		It("is included with errors", func() {
			_, err := vm.Run(`
def foo
  bar()
end

def bar
  baz()
end

def baz
  nil + 5 # whoops!
end

baz()
`)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("baz"))
			Expect(err.Error()).To(ContainSubstring("bar"))
			Expect(err.Error()).To(ContainSubstring("foo"))
		})
	})

	Context("when an error occurs in the middle of a series of statements", func() {
		It("halts execution at the error", func() {
			_, err := vm.Run(`
foo = 1
require 'some/file/that/does/not/exist/hopefully'
foo = 0
`)
			Expect(err).To(HaveOccurred())

			value, _ := vm.Get("foo")
			Expect(value).To(Equal(builtins.NewFixnum(1, vm, vm)))
		})
	})

	Describe("defining a class", func() {
		It("adds it to the global class cache", func() {
			_, err := vm.Run(`
class Foo
end
`)

			Expect(err).ToNot(HaveOccurred())
			_, err = vm.GetClass("Foo")
			Expect(err).ToNot(HaveOccurred())
		})

		It("allows a user to construct an instance of the class", func() {
			_, err := vm.Run(`
class Foo
end
`)

			fooClass := vm.MustGetClass("Foo")
			method, err := fooClass.Method("new")
			Expect(err).ToNot(HaveOccurred())

			instance, err := method.Execute(fooClass, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(instance.Class()).To(Equal(fooClass))
		})

		Context("when there are instance methods defined", func() {
			It("makes instance methods available on new instances of the class", func() {
				_, err := vm.Run(`
class Foo
  def hello
    "world"
  end
end
`)

				Expect(err).ToNot(HaveOccurred())

				fooClass, err := vm.GetClass("Foo")
				Expect(fooClass).ToNot(BeNil())

				fooInstance, err := fooClass.New(vm, vm)
				Expect(err).ToNot(HaveOccurred())
				Expect(fooInstance).ToNot(BeNil())

				method, err := fooInstance.Method("hello")
				Expect(err).ToNot(HaveOccurred())

				val, err := method.Execute(fooInstance, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(BeAssignableToTypeOf(builtins.NewString("", vm, vm)))
				Expect(val.String()).To(Equal(`"world"`))
			})
		})

		Context("when it extends a module", func() {
			It("makes the modules methods available on itself", func() {
				_, err := vm.Run(`
module Foo
  def publication
    'Chichimec-lipochrome'
  end
end

class Bar
  extend Foo
end
`)

				Expect(err).ToNot(HaveOccurred())

				barClass := vm.MustGetClass("Bar")
				Expect(barClass).To(HaveMethod("publication"))
			})
		})

		Context("when it includes a module", func() {
			It("makes the modules methods available on its instances", func() {
				_, err := vm.Run(`
module Foo
  def superinquisitive
    "tumescent-wasty"
  end
end

class Bar
  include Foo
end
`)
				Expect(err).ToNot(HaveOccurred())

				barClass := vm.MustGetClass("Bar")
				bar, err := barClass.New(vm, vm)
				Expect(err).ToNot(HaveOccurred())

				method, err := bar.Method("superinquisitive")
				Expect(err).ToNot(HaveOccurred())

				val, err := method.Execute(bar, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(val.String()).To(Equal(`"tumescent-wasty"`))
			})
		})
	})

	Describe("the ternary operator", func() {
		It("picks the first value when it is truthy", func() {
			val, err := vm.Run("foo = true ? 'a' : 'b'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal(`"a"`))
		})

		It("picks the second value when the first is falsy", func() {
			val, err := vm.Run("foo = nil ? 'a' : 'b'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal(`"b"`))
		})
	})

	Describe("eigenclasses", func() {
		BeforeEach(func() {
			_, err := vm.Run(`
object = Object.new

def object.whatever
end
`)

			Expect(err).ToNot(HaveOccurred())
		})

		It("can store instance methods for a given object", func() {
			object := vm.MustGet("object")
			Expect(object).ToNot(BeNil())
			Expect(object).To(HaveMethod("whatever"))
		})

		It("lists its methods when Kernel#singleton_methods is invoked", func() {
			list, err := vm.Run(`
object.singleton_methods
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(list.(*builtins.Array).Members()).To(ContainElement(vm.Symbols()["whatever"]))
		})
	})

	Describe("equality", func() {
		Context("with the == operator", func() {
			It("treats objects as equal when they have the same value", func() {
				result, err := vm.Run("'foo' == 'foo'")

				trueValue := vm.SingletonWithName("true")
				falseValue := vm.SingletonWithName("false")

				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(trueValue))
				result, err = vm.Run("'foo' == 'bar'")

				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(falseValue))

				result, err = vm.Run(":foo == :foo")
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(trueValue))
			})
		})
	})
})
