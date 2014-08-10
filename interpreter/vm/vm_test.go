package vm_test

import (
	"github.com/grubby/grubby/interpreter/vm/builtins"
	"github.com/grubby/grubby/interpreter/vm/builtins/errors"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/testhelpers"
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
			_, err := vm.Run(`
def foo
end`)

			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a private method on Kernel", func() {
			kernel, err := vm.Get("Kernel")
			Expect(err).ToNot(HaveOccurred())
			Expect(len(kernel.PrivateMethods())).To(Equal(1))
			Expect(kernel.PrivateMethods()[0].Name()).To(Equal("foo"))
		})

		// FIXME: to fix this, methods should know about
		// * super class methods (ideally the entire chain)
		// * included modules
		// (((   and Object needs to "include" Kernel   )))
		PIt("is also available in Object", func() {
			object, err := vm.Get("Object")
			Expect(err).ToNot(HaveOccurred())
			Expect(object.PrivateMethods()).To(ContainElement(
				builtins.NewMethod("foo", nil),
			))
		})
	})

	Describe("strings", func() {
		It("returns a ruby String object", func() {
			val, err := vm.Run("'nonrestricted-consonantize'")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewString("")))
			Expect(val.String()).To(Equal("'nonrestricted-consonantize'"))
		})
	})

	Describe("numbers", func() {
		It("returns a ruby Fixnum object", func() {
			val, err := vm.Run("5")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(builtins.NewInt(0)))
			Expect(val.String()).To(Equal("5"))
		})
	})

	Describe("the require method on Kernel", func() {
		It("searches for a file with the given name", func() {
			_, err := vm.Run("require 'something'")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(errors.NewLoadError("")))
		})

		Context("with a load path and a file to require", func() {
			BeforeEach(func() {
				SetupLoadPathWithAFileToRequire(vm)
			})

			It("requires the file", func() {
				_, err := vm.Run("require 'foo'")

				Expect(err).ToNot(HaveOccurred())

				kernel, err := vm.Get("Kernel")
				Expect(err).ToNot(HaveOccurred())
				Expect(len(kernel.PrivateMethods())).To(Equal(1))
				Expect(kernel.PrivateMethods()[0].Name()).To(Equal("foo"))
			})
		})
	})
})
