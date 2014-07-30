package vm_test

import (
	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VM", func() {
	var vm VM

	BeforeEach(func() {
		vm = NewVM()
	})

	It("has a global Object", func() {
		Expect(vm.MustGet("Object")).ToNot(BeNil())
	})

	Describe("creating a simple function", func() {
		BeforeEach(func() {
			err := vm.Run(`
def foo
end`)

			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a private method on Kernel", func() {
			kernel, err := vm.Get("Kernel")
			Expect(err).ToNot(HaveOccurred())
			Expect(kernel.PrivateMethods()).To(ContainElement(
				builtins.NewMethod("foo"),
			))
		})

		// FIXME: to fix this, methods should know about
		// * super class methods (ideally the entire chain)
		// * included modules
		// (((   and Object needs to "include" Kernel   )))
		PIt("is also available in Object", func() {
			object, err := vm.Get("Object")
			Expect(err).ToNot(HaveOccurred())
			Expect(object.PrivateMethods()).To(ContainElement(
				builtins.NewMethod("foo"),
			))
		})
	})
})
