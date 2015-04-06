package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("regular expressions", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be created with /^syntax$/", func() {
		value, err := vm.Run(`/foo.*bar/`)

		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeAssignableToTypeOf(&Regexp{}))
	})

	It("has a #match method", func() {
		result, err := vm.Run(`/foo.*bar/.match("foo WOAH bar")`)

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(vm.SingletonWithName("true")))
	})

	It("can be used to quote strings for use as regular expressions", func() {
		value, err := vm.Run("Regexp.quote('foo ^ bar')")
		Expect(err).ToNot(HaveOccurred())
		Expect(value.(*StringValue).String()).To(ContainSubstring("foo\\ \\^\\ bar"))
	})
})
