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

	Describe("defined outside of a class", func() {
		BeforeEach(func() {
			_, err := vm.Run(`
def foo
end
`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be defined on 'main'", func() {
			Expect(vm.MustGet("Kernel")).To(HavePrivateMethod("foo"))
		})

		It("should be invokeable", func() {
			_, err := vm.Run("foo()")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	It("has a reference to self", func() {
		_, err := vm.Run(`
class Foo
end
`)

		Expect(err).ToNot(HaveOccurred())

		foo, err := vm.MustGetClass("Foo").New(vm)
		Expect(err).ToNot(HaveOccurred())

		var capturedSelf Value
		foo.AddMethod(NewNativeMethod("fasciculated_stripe", vm, func(self Value, block Block, args ...Value) (Value, error) {
			capturedSelf = self
			return nil, nil
		}))

		method := foo.Method("fasciculated_stripe")
		Expect(method).ToNot(BeNil())
		method.Execute(foo, nil)

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

		Describe("named arguments", func() {
			var object Value

			Context("provided a value", func() {
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

			Context("not provided a value", func() {
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
					Expect(object.GetInstanceVariable("bar")).To(Equal(NewFixnum(1, vm)))
				})
			})
		})
	})

	Describe("return values", func() {
		var (
			result Value
			err    error
		)

		Context("when there is an explicit return expression", func() {
			BeforeEach(func() {
				result, err = vm.Run(`
def test
	if true
	  return 'hello'
	else
	  return 'world'
  end
end

test()
`)
			})

			It("returns the value provided", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(result.(*StringValue).RawString()).To(Equal("hello"))
			})
		})

		Context("when there is no explicit return expression", func() {
			BeforeEach(func() {
				result, err = vm.Run(`
def test
  'hello'
	'world'
end

test()
`)
			})

			It("returns the last value in the body of the method", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(result.(*StringValue).RawString()).To(Equal("world"))
			})
		})
	})
})
