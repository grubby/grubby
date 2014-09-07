package vm_test

import (
	"os"
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
		vm = NewVM("fake-irb-under-test")
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
			kernel := vm.MustGet("Kernel")
			method, err := kernel.PrivateMethod("foo")

			Expect(err).ToNot(HaveOccurred())
			Expect(method.Name()).To(Equal("foo"))
		})

		// FIXME: to fix this, methods should know about
		// * super class methods (ideally the entire chain)
		// * included modules
		// (((   and Object needs to "include" Kernel   )))
		PIt("is also available in Object", func() {
			object, err := vm.Get("Object")
			Expect(err).ToNot(HaveOccurred())
			Expect(object.PrivateMethods()).To(ContainElement(
				builtins.NewMethod("foo", nil),
			))
		})
	})

	Describe("strings", func() {
		It("returns a ruby String object", func() {
			val, err := vm.Run("'nonrestricted-consonantize'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewString("")))
			Expect(val.String()).To(Equal("nonrestricted-consonantize"))
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
				vm.Set("foo", builtins.NewString("superinquisitive-edacity"))
			})

			It("returns the variable referenced", func() {
				value, err := vm.Run("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(value).To(Equal(builtins.NewString("superinquisitive-edacity")))
			})
		})

		Context("when it has not been defined", func() {
			It("returns an error", func() {
				_, err := vm.Run("grigri_incircumscription")

				Expect(err).To(HaveOccurred())
			})
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
			Expect(err).To(BeAssignableToTypeOf(builtins.NewLoadError("")))
		})

		Context("with a load path and a file to require", func() {
			BeforeEach(func() {
				SetupLoadPathWithAFileToRequire(vm)
			})

			It("requires the file", func() {
				_, err := vm.Run("require 'foo'")
				Expect(err).ToNot(HaveOccurred())

				kernel := vm.MustGet("Kernel")
				method, err := kernel.PrivateMethod("foo")

				Expect(err).ToNot(HaveOccurred())
				Expect(method.Name()).To(Equal("foo"))
			})
		})
	})

	Describe("the load path", func() {
		It("is represented by $LOAD_PATH and $:", func() {
			path := vm.MustGet("LOAD_PATH")
			str := builtins.NewString("foo")
			path.(*builtins.Array).Append(str)

			Expect(vm.MustGet(":").(*builtins.Array).Members()).To(ContainElement(str))
		})
	})

	Describe("File class", func() {
		It("has a reasonable .expand_path method", func() {
			fileClass, err := vm.Get("File")
			Expect(err).ToNot(HaveOccurred())

			method, err := fileClass.Method("expand_path")
			Expect(err).ToNot(HaveOccurred())

			result, err := method.Execute(builtins.NewString("~/foobar"))
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
			Expect(value).To(Equal(builtins.NewString("albitite-compotor")))
		})
	})

	Describe("special global variables", func() {
		Describe("__FILE__", func() {
			It("inherits the name given to the vm initially", func() {
				value, err := vm.Run("__FILE__")

				Expect(err).ToNot(HaveOccurred())
				Expect(value).To(Equal(builtins.NewString("fake-irb-under-test")))
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

		Describe("ARGV", func() {
			It("has a shift method", func() {
				value, err := vm.Run("ARGV.shift")

				Expect(err).ToNot(HaveOccurred())
				Expect(value).To(Equal(builtins.Nil()))
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

				trueValue := builtins.NewTrueClass().(builtins.Class).New()

				Expect(err).ToNot(HaveOccurred())
				Expect(vm.MustGet("foo")).To(Equal(trueValue))
				Expect(vm.MustGet("bar")).To(Equal(trueValue))
			})
		})

		Describe("calling a method that does not exist", func() {
			It("raises a NoMethodError", func() {
				_, err := vm.Run("$foo.bar()")
				Expect(err).To(BeAssignableToTypeOf(builtins.NewNoMethodError("", "", "")))
			})
		})
	})
})
