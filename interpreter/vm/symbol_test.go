package vm_test

import (
	"os"
	"path/filepath"
	"reflect"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("symbols", func() {
	var (
		vm  VM
		val Value
		err error
	)

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	BeforeEach(func() {
		val, err = vm.Run(":foo")
		Expect(err).ToNot(HaveOccurred())
	})

	It("is an instance of the Symbol class", func() {
		Expect(val.Class()).To(Equal(vm.MustGetClass("Symbol")))
	})

	It("is a ruby Symbol object", func() {
		Expect(val).To(Equal(vm.Symbols()["foo"]))
	})

	It("is registered once, globally", func() {
		Expect(val).To(Equal(vm.Symbols()["foo"]))
	})

	It("is a singleton", func() {
		sameSymbol, err := vm.Run(":foo")
		Expect(err).ToNot(HaveOccurred())

		firstPointer := reflect.ValueOf(val).Pointer()
		secondPointer := reflect.ValueOf(sameSymbol).Pointer()
		Expect(secondPointer).To(Equal(firstPointer))
	})
})
