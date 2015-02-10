package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Enumerable collections", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Describe("select", func() {
		It("filters the collection given the block provided", func() {
			value, err := vm.Run("[1,2,3].select { |o| false }")
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*Array).Members()).To(BeEmpty())
		})

		It("only keeps the elements for which the block yields a truthy value", func() {
			value, err := vm.Run("[1,2,3].select { |o| o.even? }")
			Expect(err).ToNot(HaveOccurred())
			Expect(len(value.(*Array).Members())).To(Equal(1))
			Expect(value.(*Array).Members()).To(ContainElement(NewFixnum(2, vm, vm)))
		})
	})
})
