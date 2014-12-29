package vm_test

import (
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/grubby/grubby/interpreter/vm"
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

		_, ok := value.(*builtins.Array)
		Expect(ok).To(BeTrue())
	})
})
