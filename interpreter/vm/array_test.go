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
			Expect(array.Members()).To(ContainElement(vm.Symbols()["hello"]))
		})
	})

	Describe("joining the elements of an array together", func() {
		It("returns the elements joined with the given string separator", func() {
			value, err := vm.Run("[1,2,3].join(':')")
			Expect(err).ToNot(HaveOccurred())

			asStr, ok := value.(*StringValue)
			Expect(ok).To(BeTrue())
			Expect(asStr.RawString()).To(Equal("1:2:3"))
		})
	})

	Describe("mapping over the items in an array", func() {
		var (
			value Value
			err   error
		)

		BeforeEach(func() {
			value, err = vm.Run("['1', '2', '3'].map(&:to_i)")
		})

		It("runs the provided block over each of the items, returning the new collection", func() {
			Expect(err).ToNot(HaveOccurred())

			array, ok := value.(*Array)
			Expect(ok).To(BeTrue())
			Expect(len(array.Members())).To(Equal(3))
			Expect(array.Members()).To(ContainElement(NewFixnum(1, vm)))
			Expect(array.Members()).To(ContainElement(NewFixnum(2, vm)))
			Expect(array.Members()).To(ContainElement(NewFixnum(3, vm)))
		})
	})

	Describe("iterating over the items in an array", func() {
		It("calls the provided block once with each item in the array", func() {
			value, err := vm.Run(`
array = []
[1,2,3].each {|i| array.unshift(i) }
array
`)

			Expect(err).ToNot(HaveOccurred())

			array, ok := value.(*Array)
			Expect(ok).To(BeTrue())

			Expect(len(array.Members())).To(Equal(3))
			Expect(array.Members()).To(ContainElement(NewFixnum(1, vm)))
			Expect(array.Members()).To(ContainElement(NewFixnum(2, vm)))
			Expect(array.Members()).To(ContainElement(NewFixnum(3, vm)))
		})
	})

	Describe("#any?", func() {
		It("takes the given block, calling it with each member until the block returns a truthy value", func() {
			value, err := vm.Run(`
visited = []
[nil,false,3,4].any? { |i| visited.unshift(i); i }
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(vm.SingletonWithName("true")))

			visited := vm.MustGet("visited").(*Array).Members()
			Expect(len(visited)).To(Equal(3))
			Expect(visited[0].String()).To(ContainSubstring("3")) // unshift is unkind
			Expect(visited[1].String()).To(ContainSubstring("false"))
			Expect(visited[2].String()).To(ContainSubstring(""))
		})
	})
})
