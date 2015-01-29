package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arrays", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be constructed as an array literal", func() {
		value, err := vm.Run("[:hello, :world]")
		Expect(err).ToNot(HaveOccurred())
		Expect(value).ToNot(BeNil())

		_, ok := value.(*Array)
		Expect(ok).To(BeTrue())
	})

	Describe("subtracting one array from another", func() {
		It("returns the elements in the first that are not in the latter", func() {
			value, err := vm.Run("[:hello, :world] - [:cruel, :world]")
			Expect(err).ToNot(HaveOccurred())

			array, ok := value.(*Array)
			Expect(ok).To(BeTrue())
			Expect(len(array.Members())).To(Equal(1))
			Expect(array.Members()).To(ContainElement(NewSymbol("hello", vm)))
		})
	})
})
