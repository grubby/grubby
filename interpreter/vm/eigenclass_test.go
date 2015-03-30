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

var _ = Describe("Eigenclasses", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Describe("opened by defining a method on a single object", func() {
		BeforeEach(func() {
			_, err := vm.Run(`
object = Object.new

def object.whatever
end
`)

			Expect(err).ToNot(HaveOccurred())
		})

		It("can store instance methods for a given object", func() {
			object := vm.MustGet("object")
			Expect(object).ToNot(BeNil())
			Expect(object).To(HaveMethod("whatever"))
		})

		It("lists its methods when Kernel#singleton_methods is invoked", func() {
			list, err := vm.Run(`
object.singleton_methods
`)

			Expect(err).ToNot(HaveOccurred())

			listAsArray, ok := list.(*Array)
			Expect(ok).To(BeTrue())

			Expect(listAsArray.Members()).To(ContainElement(vm.Symbols()["whatever"]))
			Expect(listAsArray.Members()).ToNot(ContainElement(vm.Symbols()["puts"]))
		})
	})

	Describe("opened by the 'class << self' syntax", func() {
		BeforeEach(func() {
			_, err := vm.Run(`
class MyObject
  class << self
    def my_method
      'bewildering-wildebeest'
    end
  end
end
`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("can store instance methods", func() {
			class := vm.MustGetClass("MyObject")
			Expect(class).To(HaveMethod("my_method"))
		})
	})
})
