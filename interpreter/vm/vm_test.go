package vm_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
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
			method, err := kernel.Method("foo")

			Expect(err).ToNot(HaveOccurred())
			Expect(method.Name()).To(Equal("foo"))
			Expect(method.IsPrivate()).To(Equal(true))
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

	Describe("nil", func() {
		var (
			val Value
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

	Describe("blocks", func() {
		It("can be used to close over variables", func() {
			value, err := vm.Run(`
def test(needle)
  [1,2,3,4,5].select {|o| o == needle}
end

test(5)
`)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(value.(*Array).Members())).To(Equal(1))
			Expect(value.(*Array).Members()).To(ContainElement(NewFixnum(5, vm, vm)))
		})
	})

	Describe("a reference to a variable", func() {
		Context("after it has been defined", func() {
			BeforeEach(func() {
				vm.Set("foo", NewString("superinquisitive-edacity", vm, vm))
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
			Expect(err).ToNot(BeAssignableToTypeOf(NewLoadError("", "")))
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
			Expect(err).To(BeAssignableToTypeOf(NewLoadError("", "")))
		})

		Context("with a load path and a file to require", func() {
			BeforeEach(func() {
				SetupLoadPathWithAFileToRequire(vm)
			})

			It("requires the file", func() {
				_, err := vm.Run("require 'foo'")
				Expect(err).ToNot(HaveOccurred())

				kernel := vm.Modules()["Kernel"]
				method, err := kernel.Method("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(method.IsPrivate()).To(Equal(true))
				Expect(method.Name()).To(Equal("foo"))
			})
		})
	})

	Describe("the load path", func() {
		It("is represented by $LOAD_PATH and $:", func() {
			path := vm.MustGet("LOAD_PATH")
			str := NewString("foo", vm, vm)
			path.(*Array).Append(str)

			Expect(vm.MustGet(":").(*Array).Members()).To(ContainElement(str))
		})
	})

	Describe("File class", func() {
		It("has a reasonable .expand_path method", func() {
			fileClass := vm.ClassWithName("File")
			Expect(fileClass).ToNot(BeNil())

			method, err := fileClass.Method("expand_path")
			Expect(err).ToNot(HaveOccurred())

			result, err := method.Execute(fileClass, nil, NewString("~/foobar", vm, vm))
			Expect(err).ToNot(HaveOccurred())

			expectedPath := fmt.Sprintf("%s/%s", os.Getenv("HOME"), "foobar")
			Expect(result.String()).To(Equal(expectedPath))
		})
	})

	Describe("assignment to a variable", func() {
		It("stores the value assigned", func() {
			_, err := vm.Run("foo = 'albitite-compotor'")
			Expect(err).ToNot(HaveOccurred())

			value, err := vm.Get("foo")
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("albitite-compotor"))
		})
	})

	Describe("constants", func() {
		It("can be nested inside of a module", func() {
			_, err := vm.Run("Kernel::Foobar = 1")
			Expect(err).ToNot(HaveOccurred())

			kernel := vm.MustGet("Kernel").(Module)
			Expect(kernel.Constant("Foobar")).To(Equal(NewFixnum(1, vm, vm)))
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

		Describe("__LINE__", func() {
			It("is equal to the line number it appeared at", func() {
				result, err := vm.Run(`
__LINE__
`)

				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(NewFixnum(1, vm, vm)))
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
			Expect(err).To(BeAssignableToTypeOf(NewNoMethodError("", "", "", "")))
			Expect(err.Error()).To(ContainSubstring("undefined method 'world' for \"hello\":String"))
		})
	})

	XDescribe("stack traces", func() {
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

	Describe("when an error occurs", func() {
		Context("in the middle of some statements", func() {
			It("halts execution at the error", func() {
				_, err := vm.Run(`
foo = 1
require 'some/file/that/does/not/exist/hopefully'
foo = 0
`)
				Expect(err).To(HaveOccurred())

				value, _ := vm.Get("foo")
				Expect(value).To(Equal(NewFixnum(1, vm, vm)))
			})
		})

		Context("inside of a module declaration", func() {
			BeforeEach(func() {
				_, err := vm.Run(`
counter = 1

module Foo
  counter = 2
  foo # this should fail, no such method
  counter = 3
end

counter = 4
`)
				Expect(err).To(HaveOccurred())
			})

			It("halts the module declaration", func() {
				Expect(vm.MustGet("counter")).To(Equal(NewFixnum(2, vm, vm)))
			})
		})
	})

	Describe("the ternary operator", func() {
		It("picks the first value when it is truthy", func() {
			val, err := vm.Run("foo = true ? 'a' : 'b'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal("a"))
		})

		It("picks the second value when the first is falsy", func() {
			val, err := vm.Run("foo = nil ? 'a' : 'b'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal("b"))
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

	Describe("self", func() {
		var err error
		var output string

		JustBeforeEach(func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("outside of an instance method", func() {
			BeforeEach(func() {
				output = SwapStdout(func() {
					_, err = vm.Run("puts self")
				})
			})

			It("should be main", func() {
				Expect(output).To(ContainSubstring("main"))
			})
		})

		Context("inside an instance method", func() {
			BeforeEach(func() {
				output = SwapStdout(func() {
					_, err = vm.Run(`
class Foo
  def inspect
    self
  end
end

puts Foo.new.inspect
`)
				})
			})

			It("should refer to the instance itself", func() {
				Expect(output).To(ContainSubstring("<Foo:0x"))
			})
		})
	})
})
