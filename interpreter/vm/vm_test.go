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
		grubbyHome := filepath.Join(os.Getenv("HOME"), ".grubby")
		vm = NewVM(grubbyHome, "fake-irb-under-test")
	})

	Describe("Object", func() {
		It("is a Class", func() {
			o := vm.MustGet("Object")
			Expect(o.Class().String()).To(Equal("Class"))
		})

		It("inherits from BasicObject", func() {
			Expect(vm.MustGet("Object").(Class).SuperClass().String()).To(Equal("BasicObject"))
		})

		Describe("#object_id", func() {
			It("is available to call", func() {
				_, err := vm.Run("Object.new.object_id")
				Expect(err).ToNot(HaveOccurred())
			})

			It("can be aliased", func() {
				_, err := vm.Run(`
class Object
  alias_method :my_object_id, :object_id
end
`)
				Expect(err).ToNot(HaveOccurred())
			})
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
			method := kernel.Method("foo")

			Expect(method).ToNot(BeNil())
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
			Expect(value.(*Array).Members()).To(ContainElement(NewFixnum(5, vm)))
		})

		It("can handle string interpolation", func() {
			value, err := vm.Run(`[1,2,3].map {|o| "foo #{o} bar" }`)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(value.(*Array).Members())).To(Equal(3))
			Expect(value.(*Array).Members()[0].(*StringValue).RawString()).To(ContainSubstring("foo 1 bar"))
		})
	})

	Describe("a reference to a variable", func() {
		Context("after it has been defined", func() {
			BeforeEach(func() {
				vm.Set("foo", NewString("superinquisitive-edacity", vm))
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

	Describe("the standard lib", func() {
		It("is available to require", func() {
			_, err := vm.Run("require 'fileutils'")
			Expect(err).ToNot(HaveOccurred())
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
				method := kernel.Method("foo")

				Expect(method).ToNot(BeNil())
				Expect(method.IsPrivate()).To(Equal(true))
				Expect(method.Name()).To(Equal("foo"))
			})
		})

		Describe("with the same file twice", func() {
			BeforeEach(func() {
				SetupLoadPathWithAFileToRequire(vm)
			})

			It("does not require the same file twice", func() {
				value, err := vm.Run(`
require 'foo'

FOO_CONST = 'bar'

require 'foo'
`)
				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).ToNot(ContainSubstring("bar"))
			})
		})
	})

	Describe("the load path", func() {
		It("is represented by $LOAD_PATH and $:", func() {
			path := vm.MustGet("LOAD_PATH")
			str := NewString("foo", vm)
			path.(*Array).Append(str)

			Expect(vm.MustGet(":").(*Array).Members()).To(ContainElement(str))
		})
	})

	Describe("global variables", func() {
		It("defaults to nil if it has not been defined yet", func() {
			value, err := vm.Run("$garbage")
			Expect(err).ToNot(HaveOccurred())

			nilInstance := vm.SingletonWithName("nil")
			Expect(value).To(Equal(nilInstance))
		})
	})

	Describe("the File class", func() {
		It("has a reasonable .expand_path method", func() {
			fileClass := vm.ClassWithName("File")
			Expect(fileClass).ToNot(BeNil())

			method := fileClass.Method("expand_path")
			Expect(method).ToNot(BeNil())

			result, err := method.Execute(fileClass, nil, NewString("~/foobar", vm))
			Expect(err).ToNot(HaveOccurred())

			expectedPath := fmt.Sprintf("%s/%s", os.Getenv("HOME"), "foobar")
			Expect(result.String()).To(Equal(expectedPath))
		})
	})

	Describe("the Dir class", func() {
		Describe(".pwd", func() {
			It("returns the current working directory of the process", func() {
				value, err := vm.Run("Dir.pwd")
				Expect(err).ToNot(HaveOccurred())

				dir, ok := value.(*StringValue)
				Expect(ok).To(BeTrue())

				cwd, _ := os.Getwd()

				Expect(dir.RawString()).To(Equal(cwd))
			})
		})
	})

	Describe("the Process class", func() {
		Describe(".pid", func() {
			It("returns the id of the current process", func() {
				value, err := vm.Run("Process.pid")
				Expect(err).ToNot(HaveOccurred())

				Expect(value.String()).To(Equal(fmt.Sprintf("%d", os.Getpid())))
			})
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
			value, err := vm.Run("Kernel::Foobar = 1; Kernel::Foobar")
			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(NewFixnum(1, vm)))

			kernel := vm.MustGet("Kernel").(Module)
			Expect(kernel.Constant("Foobar")).To(Equal(NewFixnum(1, vm)))
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
				Expect(result).To(Equal(NewFixnum(1, vm)))
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
				Expect(value).To(Equal(NewFixnum(1, vm)))
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
				Expect(vm.MustGet("counter")).To(Equal(NewFixnum(2, vm)))
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

	Describe("truthiness", func() {
		It("considers nil to be falsey", func() {
			value, err := vm.Run(`
value = 'okay'
if nil
  value = 'whack'
end

value
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(ContainSubstring("okay"))
		})

		It("considers negated nil to be truthy", func() {
			value, err := vm.Run("!nil")
			Expect(err).ToNot(HaveOccurred())

			trueValue := vm.SingletonWithName("true")
			Expect(value).To(Equal(trueValue))
		})
	})

	Describe("Kernel#at_exit", func() {
		It("can be triggered by calling vm.exit", func() {
			_, err := vm.Run(`
value = 'too soon'

at_exit do
  value = 'just right'
end
`)
			Expect(err).ToNot(HaveOccurred())

			vm.Exit()
			Expect(vm.MustGet("value").String()).To(ContainSubstring("just right"))
		})
	})

	Describe("opening up a class again", func() {
		Context("with multiple methods declared in each block", func() {
			It("preserves the earlier definition of the class", func() {
				value, err := vm.Run(`
class Foo
  def initial_method
  end
end

class Foo
  def monkey_patched_method
  end
end
`)
				Expect(err).ToNot(HaveOccurred())
				Expect(value.(Class)).To(HaveInstanceMethod("initial_method"))
				Expect(value.(Class)).To(HaveInstanceMethod("monkey_patched_method"))
			})
		})

		Context("and calling the new methods from main", func() {
			It("does not fail", func() {
				value, err := vm.Run(`

class Object
  def sup
    'not much, just monkey patching'
  end
end

sup`)

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(ContainSubstring("not much, just monkey patching"))
			})
		})

		Context("and calling the new methods on an existing class", func() {
			It("does not fail", func() {
				value, err := vm.Run(`

class Object
  def sup
    'not much, just monkey patching'
  end
end

File.sup`)

				Expect(err).ToNot(HaveOccurred())
				Expect(value.String()).To(ContainSubstring("not much, just monkey patching"))
			})
		})
	})

	Describe("boolean 'and'", func() {
		It("can be used to check if two values are truthy", func() {
			value, err := vm.Run("true and 'a'")

			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(ContainSubstring("a"))
		})
	})

	Describe("the defined? keyword", func() {
		It("can be used to check if a value is defined", func() {
			value, err := vm.Run("defined? a")
			Expect(err).ToNot(HaveOccurred())

			nilInstance := vm.SingletonWithName("nil")
			Expect(value).To(Equal(nilInstance))
		})
	})
})
