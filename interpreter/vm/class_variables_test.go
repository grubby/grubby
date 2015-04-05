package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("class variables", func() {
	var (
		vm          VM
		fooClass    Class
		fooInstance Value
	)

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
		vm.Run(`
class Foo
  def initialize
    @@foo = 'hello'
  end

  def test
    @@foo
  end

  def test=(something)
    @@foo = something
  end
end
`)

		fooClass = vm.ClassWithName("Foo")
		Expect(fooClass).ToNot(BeNil())

		fooInstance, err = fooClass.New(vm)
		Expect(err).ToNot(HaveOccurred())
	})

	It("can be used to read values", func() {
		foo := fooInstance.GetClassVariable("foo")
		Expect(foo.(*StringValue).RawString()).To(Equal("hello"))
	})

	It("can be used to write values", func() {
		anotherFoo, err := fooClass.New(vm)
		Expect(err).ToNot(HaveOccurred())

		method := anotherFoo.Method("test=")
		Expect(method).ToNot(BeNil())

		_, err = method.Execute(anotherFoo, nil, NewString("world", vm))
		Expect(err).ToNot(HaveOccurred())
		Expect(fooInstance.GetClassVariable("foo")).To(Equal(anotherFoo.GetClassVariable("foo")))
	})
})
