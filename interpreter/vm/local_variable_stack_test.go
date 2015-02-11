package vm_test

import (
	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The local variable stack", func() {
	var subject *LocalVariableStack
	var vm VM

	BeforeEach(func() {
		vm = NewVM("fake-rubyhome", "fake-irb-under-test")
		subject = NewLocalVariableStack()
		subject.Unshift()
	})

	Describe("stores values", func() {
		var storedValue Value
		BeforeEach(func() {
			storedValue = NewFixnum(5, vm, vm)
			subject.Store("foo", storedValue)
		})

		It("that can be retrieved later", func() {
			Expect(subject.Retrieve("foo")).To(Equal(storedValue))
		})

		Context("but when a new frame is added to the top of the stack", func() {
			BeforeEach(func() {
				subject.Unshift()
			})

			It("does not return the value that was originally stored", func() {
				_, err := subject.Retrieve("foo")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
