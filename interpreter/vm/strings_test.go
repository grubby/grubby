package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be constructed from a string literal", func() {
		val, err := vm.Run("'nonrestricted-consonantize'")

		Expect(err).ToNot(HaveOccurred())
		Expect(val).To(BeAssignableToTypeOf(NewString("", vm)))
		Expect(val.String()).To(Equal("nonrestricted-consonantize"))
	})

	It("has a + method", func() {
		val, err := vm.Run("'foo' + 'bar'")

		Expect(err).ToNot(HaveOccurred())
		Expect(val.String()).To(Equal(NewString("foobar", vm).String()))
	})

	Describe("interpolating values inside of a string", func() {
		It("can be done with double quoted strings", func() {
			value, err := vm.Run(`
greetings = 'hello'
adj       = 'cruel'
place     = 'world'
"#{greetings} #{adj} #{place}"
`)
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("hello cruel world"))
		})

		It("calls to_s on the values being interpolated", func() {
			value, err := vm.Run(`"whoops all #{:crunchberries}"`)
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("whoops all crunchberries"))
		})

		It("cannot be done with single quoted strings", func() {
			value, err := vm.Run(`
adj = 'cruel'
'hello #{adj} world'
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("hello #{adj} world"))
		})
	})

	Describe("concatenating strings", func() {
		It("can be done with the shovel operator", func() {
			value, err := vm.Run("'hello' << ' world'")
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("hello world"))
		})
	})

	Describe("#freeze", func() {
		It("can be used to turn a string into an immutable string", func() {
			_, err := vm.Run("str = 'hello'.freeze; str << 'world'")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("RuntimeError: can't modify frozen String"))
		})
	})

	Describe("#intern", func() {
		It("can be used to create a symbol with string's value as its name", func() {
			value, err := vm.Run("'hello world'.intern")
			Expect(err).ToNot(HaveOccurred())

			Expect(value).ToNot(BeNil())
			Expect(value).To(BeAssignableToTypeOf(&SymbolValue{}))
			Expect(value).To(Equal(vm.SymbolWithName("hello world"))) // singleton instance
		})
	})

	Describe("#split", func() {
		It("splits on the given separator", func() {
			result, err := vm.Run("'hello world'.split(' ')")
			Expect(err).ToNot(HaveOccurred())

			array, ok := result.(*Array)
			Expect(ok).To(BeTrue())
			Expect(len(array.Members())).To(Equal(2))
			Expect(array.Members()[0].String()).To(Equal("hello"))
			Expect(array.Members()[1].String()).To(Equal("world"))
		})
	})

	Describe("#encode", func() {
		It("should be implmented", func() {
			result, err := vm.Run(`
str = "<hello><world/></hello>"
str.encode('ISO-8859-1', :undef => :replace, :invalid => :replace)
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(result.String()).To(ContainSubstring("<hello><world/></hello>"))
		})
	})

	Describe("#=~", func() {
		XIt("should match on the given regular expression", func() {
			_, err := vm.Run("a = hello =~ /ello/; b = hello =~ /nope/")
			Expect(err).ToNot(HaveOccurred())
			Expect(vm.MustGet("a").String()).To(ContainSubstring("1"))
			Expect(vm.MustGet("b").String()).To(Equal(vm.SingletonWithName("nil")))
		})
	})
})
