package vm_test

import (
	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The call stack", func() {
	var subject *CallStack

	BeforeEach(func() {
		subject = NewCallStack()
	})

	It("initially has an empty stack", func() {
		Expect(subject.Frames).To(BeEmpty())
	})

	Describe("adding a method to the stack", func() {
		BeforeEach(func() {
			subject.Unshift("some-method", "some-file", 0)
		})

		It("should now have one item in its frames", func() {
			Expect(len(subject.Frames)).To(Equal(1))
			Expect(subject.Frames[0].File).To(Equal("some-file"))
			Expect(subject.Frames[0].Method).To(Equal("some-method"))
		})

		Context("when that frame is removed", func() {
			BeforeEach(func() {
				subject.Shift()
			})

			It("then has no frames again", func() {
				Expect(subject.Frames).To(BeEmpty())
			})
		})

		Context("when more frames are added", func() {
			BeforeEach(func() {
				subject.Unshift("method2", "file2", 1)
				subject.Unshift("method3", "file3", 2)
			})

			It("should add the last frame to the front of the stack", func() {
				Expect(len(subject.Frames)).To(Equal(3))
				Expect(subject.Frames[0].File).To(Equal("file3"))
				Expect(subject.Frames[0].Method).To(Equal("method3"))
			})
		})
	})
})
