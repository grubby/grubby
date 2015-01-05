package vm_test

import (
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("instance variables", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be used to store values", func() {
		vm.Run(`
class Foo
  def initialize
    @foo = :bar
  end
end
`)

		fooClass := vm.ClassWithName("Foo")
		Expect(fooClass).ToNot(BeNil())

		fooInstance, err := fooClass.New(vm)

		Expect(err).ToNot(HaveOccurred())
		Expect(fooInstance.GetInstanceVariable("foo")).To(Equal(builtins.NewSymbol("bar", vm)))
	})
})
