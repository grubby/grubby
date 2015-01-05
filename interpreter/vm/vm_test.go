package vm_test

import (
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
			Expect(val).To(BeAssignableToTypeOf(builtins.NewString("", vm)))
			Expect(val.String()).To(Equal("nonrestricted-consonantize"))
		})

		It("has a + method", func() {
			val, err := vm.Run("'foo' + 'bar'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal(builtins.NewString("foobar", vm).String()))
		})
	})

	Describe("interpreting a number", func() {
		It("returns a ruby Fixnum object", func() {
			val, err := vm.Run("5")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewInt(0)))
			Expect(val.String()).To(Equal("5"))
		})
	})

	Describe("interpreting a float", func() {
		It("returns a ruby Float object", func() {
			val, err := vm.Run("5.123")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewFloat(0.0)))
			Expect(val).To(Equal(builtins.NewFloat(5.123)))
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

		It("returns a ruby Symbol object", func() {
			Expect(val).To(Equal(builtins.NewSymbol("foo")))
		})

		It("registers the symbol globally", func() {
			Expect(vm.Symbols()).To(ContainElement(builtins.NewSymbol("foo")))
		})

		It("records the symbol only once", func() {
			sameSymbol, err := vm.Run(":foo")
			Expect(err).ToNot(HaveOccurred())

			firstPointer := reflect.ValueOf(val).Pointer()
			secondPointer := reflect.ValueOf(sameSymbol).Pointer()
			Expect(secondPointer).To(Equal(firstPointer))
		})
	})

	Describe("a reference to a variable", func() {
		Context("when it already is defined", func() {
			BeforeEach(func() {
				vm.Set("foo", builtins.NewString("superinquisitive-edacity", vm))
			})

			It("returns the variable referenced", func() {
				value, err := vm.Run("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal("superinquisitive-edacity"))
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
			str := builtins.NewString("foo", vm)
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

			keys, err := keysMethod.Execute(hash)
			Expect(err).ToNot(HaveOccurred())

			keyArray, ok := keys.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(keyArray.Members()).To(ContainElement(builtins.NewSymbol("key")))

			valuesMethod, err := hash.Method("values")
			Expect(err).ToNot(HaveOccurred())

			values, err := valuesMethod.Execute(hash)
			Expect(err).ToNot(HaveOccurred())

			valueArray, ok := values.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(valueArray.Members()).To(ContainElement(builtins.NewSymbol("value")))

			method, err := hash.Method("[]=")
			Expect(err).ToNot(HaveOccurred())

			_, err = method.Execute(hash, builtins.NewSymbol("foo"), builtins.NewSymbol("bar"))
			Expect(err).ToNot(HaveOccurred())

			keys, err = keysMethod.Execute(hash)
			Expect(err).ToNot(HaveOccurred())

			keyArray, ok = keys.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(keyArray.Members()).To(ContainElement(builtins.NewSymbol("foo")))

			values, err = valuesMethod.Execute(hash)
			Expect(err).ToNot(HaveOccurred())

			valueArray, ok = values.(*builtins.Array)
			Expect(ok).To(BeTrue())
			Expect(valueArray.Members()).To(ContainElement(builtins.NewSymbol("bar")))
		})
	})

	Describe("File class", func() {
		It("has a reasonable .expand_path method", func() {
			fileClass := vm.ClassWithName("File")
			Expect(fileClass).ToNot(BeNil())

			method, err := fileClass.Method("expand_path")
			Expect(err).ToNot(HaveOccurred())

			result, err := method.Execute(fileClass, builtins.NewString("~/foobar", vm))
			Expect(err).ToNot(HaveOccurred())
			Expect(result.String()).To(Equal(os.Getenv("HOME") + "/foobar"))
		})
	})

	Describe("assignment to a variable", func() {
		It("stores the value assigned", func() {
			_, err := vm.Run("foo = 'albitite-compotor'")
			Expect(err).ToNot(HaveOccurred())

			value, err := vm.Get("foo")
			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(Equal("albitite-compotor"))
		})
	})

	Describe("special global variables", func() {
		Describe("__FILE__", func() {
			It("inherits the name given to the vm initially", func() {
				value, err := vm.Run("__FILE__")

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(Equal("fake-irb-under-test"))
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

			nilInstance, err := vm.ClassWithName("Nil").New(vm)

			Expect(err).ToNot(HaveOccurred())
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

			trueValue, err := builtins.NewTrueClass(vm).(builtins.Class).New(vm)

			Expect(err).ToNot(HaveOccurred())
			Expect(vm.MustGet("foo")).To(Equal(trueValue))
			Expect(vm.MustGet("bar")).To(Equal(trueValue))
		})
	})

	Describe("calling a method that does not exist", func() {
		It("raises a NoMethodError", func() {
			_, err := vm.Run("$foo.bar()")
			Expect(err).To(BeAssignableToTypeOf(builtins.NewNoMethodError("", "", "", "")))
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
			Expect(value).To(Equal(builtins.NewInt(1)))
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

			instance, err := method.Execute(fooClass)
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

				fooInstance, err := fooClass.New(vm)
				Expect(err).ToNot(HaveOccurred())
				Expect(fooInstance).ToNot(BeNil())

				method, err := fooInstance.Method("hello")
				Expect(err).ToNot(HaveOccurred())

				val, err := method.Execute(fooInstance)
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(BeAssignableToTypeOf(builtins.NewString("", vm)))
				Expect(val.String()).To(Equal("world"))
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
				bar, err := barClass.New(vm)
				Expect(err).ToNot(HaveOccurred())

				method, err := bar.Method("superinquisitive")
				Expect(err).ToNot(HaveOccurred())

				val, err := method.Execute(bar)
				Expect(err).ToNot(HaveOccurred())
				Expect(val.String()).To(Equal("tumescent-wasty"))
			})
		})
	})
})
