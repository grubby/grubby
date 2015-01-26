package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/grubby/grubby/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("methods", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("has a reference to self", func() {
		_, err := vm.Run(`
class Foo
end
`)

		Expect(err).ToNot(HaveOccurred())

		foo, err := vm.MustGetClass("Foo").New(vm, vm)
		Expect(err).ToNot(HaveOccurred())

		var capturedSelf Value
		foo.AddMethod(NewNativeMethod("fasciculated_stripe", vm, vm, func(self Value, args ...Value) (Value, error) {
			capturedSelf = self
			return nil, nil
		}))

		m, err := foo.Method("fasciculated_stripe")
		Expect(err).ToNot(HaveOccurred())
		m.Execute(foo)

		Expect(capturedSelf).To(Equal(foo))
	})

	Context("defined by the user", func() {
		It("can refer to methods included by Kernel", func() {
			output := SwapStdout(func() {
				_, err := vm.Run(`
module Foo
  def self.greets
    puts 'hello world'
  end
end

Foo.greets
`)

				Expect(err).ToNot(HaveOccurred())
			})

			Expect(output).To(ContainSubstring("hello world"))
		})

		It("can reference its arguments by name", func() {
			output := SwapStdout(func() {
				_, err := vm.Run(`
module Foo
  def self.my_puts(arg)
    puts arg
  end
end

Foo.my_puts 'hello world'
`)

				Expect(err).ToNot(HaveOccurred())
			})

			Expect(output).To(ContainSubstring("hello world"))
		})

		Describe("default values for named arguments", func() {
			var object Value

			Context("when they are provided a value", func() {
				BeforeEach(func() {
					var err error
					object, err = vm.Run(`
class Acetylthymol
  def stroboscopical(foo, bar = 1)
    @foo = foo
    @bar = bar
  end
end

ace = Acetylthymol.new
ace.stroboscopical('abc', 'def')
ace
`)
					Expect(err).ToNot(HaveOccurred())
				})

				It("takes on the value provided", func() {
					Expect(object).ToNot(BeNil())

					ivar := object.GetInstanceVariable("foo")
					Expect(ivar.Class().String()).To(Equal("String"))
					Expect(ivar.String()).To(Equal("abc"))

					ivar = object.GetInstanceVariable("bar")
					Expect(ivar.Class().String()).To(Equal("String"))
					Expect(ivar.String()).To(Equal("def"))
				})
			})

			Context("when they are not provided a value", func() {
				BeforeEach(func() {
					var err error
					object, err = vm.Run(`
class Acetylthymol
  def stroboscopical(foo, bar = 1)
    @foo = foo
    @bar = bar
  end
end

ace = Acetylthymol.new
ace.stroboscopical('abc')
ace
`)
					Expect(err).ToNot(HaveOccurred())
				})

				It("takes on the default value", func() {
					Expect(object).ToNot(BeNil())
					Expect(object.GetInstanceVariable("foo")).To(EqualRubyString("abc"))
					Expect(object.GetInstanceVariable("bar")).To(Equal(NewFixnum(1, vm, vm)))
				})
			})
		})
	})
})
