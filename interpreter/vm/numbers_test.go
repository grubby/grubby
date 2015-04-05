package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("numbers", func() {
	var vm VM
	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Describe("integers", func() {
		It("returns a ruby Fixnum for a number literal", func() {
			val, err := vm.Run("5")

			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(BeAssignableToTypeOf(NewFixnum(0, vm)))
			Expect(val.String()).To(Equal("5"))
			Expect(val.Class()).To(Equal(vm.MustGetClass("Fixnum")))
		})

		It("treats numbers as singletons", func() {
			val, err := vm.Run("5 == 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(Equal(vm.SingletonWithName("true")))
		})

		It("has a + method", func() {
			val, err := vm.Run("11 + 31")

			Expect(err).ToNot(HaveOccurred())
			Expect(val.String()).To(Equal(NewFixnum(42, vm).String()))
		})

		It("has a #nonzero? method", func() {
			val, err := vm.Run("5.nonzero?")
			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(Equal(vm.SingletonWithName("5")))

			val, err = vm.Run("0.nonzero?")
			Expect(err).ToNot(HaveOccurred())
			Expect(val).To(Equal(vm.SingletonWithName("nil")))
		})
	})

	Describe("floats", func() {
		It("interprets floats as a ruby Float value", func() {
			val, err := vm.Run("5.123")

			Expect(err).ToNot(HaveOccurred())

			Expect(val).To(BeAssignableToTypeOf(NewFloat(0.0, vm)))
			Expect(val.Class()).To(Equal(vm.MustGetClass("Float")))

			asFloat, ok := val.(*FloatValue)
			Expect(ok).To(BeTrue())
			Expect(asFloat.ValueAsFloat()).To(Equal(5.123))
		})
	})
})
