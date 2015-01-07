package vm_test

import (
	"os"
	"path/filepath"
	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("classes", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Describe(".new", func() {
		It("returns an error when initializing the object would fail", func() {
			_, err := vm.Run(`
class Microclimatology
  def initialize
    overchildish
  end
end

Microclimatology.new
`)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("NameError"))
			Expect(err.Error()).To(ContainSubstring("undefined local variable or method 'overchildish'"))
		})
	})

	Describe("class attribute methods", func() {
		Describe(".attr_reader :symbol", func() {
			It("creates a getter and setter on instances of the class", func() {
				_, err := vm.Run(`
class Foo
  attr_reader :quaternion_vinic
end
`)

				Expect(err).ToNot(HaveOccurred())

				foo, err := vm.MustGetClass("Foo").New(vm)
				Expect(err).ToNot(HaveOccurred())

				reader, err := foo.Method("quaternion_vinic")
				Expect(err).ToNot(HaveOccurred())

				val, err := reader.Execute(foo)
				Expect(err).ToNot(HaveOccurred())

				nilInstance, err := vm.ClassWithName("Nil").New(vm)

				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal(nilInstance))
			})
		})

		Describe(".attr_writer :some_symbol", func() {
			It("creates a setter on instances of the class", func() {
				_, err := vm.Run(`
class Foo
  attr_writer :chrysobull_nonmonarchist
end
`)

				Expect(err).ToNot(HaveOccurred())

				fooClass := vm.MustGetClass("Foo")
				foo, err := fooClass.New(vm)
				Expect(err).ToNot(HaveOccurred())

				reader, err := foo.Method("chrysobull_nonmonarchist=")
				Expect(err).ToNot(HaveOccurred())

				_, err = reader.Execute(foo, NewString("lyncher-mudslinger", vm))
				Expect(err).ToNot(HaveOccurred())

				// TODO: assert on the instance variable via instance_variable_get
			})
		})

		Describe(".attr_accessor :some_symbol", func() {
			It("creates a getter and setter on instances of the class", func() {
				_, err := vm.Run(`
class Foo
  attr_accessor :pieless_bothlike
end
`)

				Expect(err).ToNot(HaveOccurred())

				fooClass := vm.MustGetClass("Foo")
				foo, err := fooClass.New(vm)
				Expect(err).ToNot(HaveOccurred())

				reader, err := foo.Method("pieless_bothlike")
				Expect(err).ToNot(HaveOccurred())

				val, err := reader.Execute(foo)
				Expect(err).ToNot(HaveOccurred())

				nilInstance, err := vm.ClassWithName("Nil").New(vm)
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal(nilInstance))

				writer, err := foo.Method("pieless_bothlike=")
				Expect(err).ToNot(HaveOccurred())

				_, err = writer.Execute(foo, NewString("unordainable-luthier", vm))
				Expect(err).ToNot(HaveOccurred())

				val, err = reader.Execute(foo)
				Expect(val.String()).To(Equal("unordainable-luthier"))
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
