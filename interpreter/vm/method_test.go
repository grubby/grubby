package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("methods", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	Context("defined by the user", func() {
		It("can refer to methods included by Kernel", func() {
			output := SwapStdout(func() {
				_, err := vm.Run(`
module Foo
  def self.greets
    puts 'hello world'
  end
end

Foo.greets
`)

				Expect(err).ToNot(HaveOccurred())
			})

			Expect(output).To(ContainSubstring("hello world"))
		})

		It("can reference its arguments by name", func() {
			output := SwapStdout(func() {
				_, err := vm.Run(`
module Foo
  def self.my_puts(arg)
    puts arg
  end
end

Foo.my_puts 'hello world'
`)

				Expect(err).ToNot(HaveOccurred())
			})

			Expect(output).To(ContainSubstring("hello world"))
		})
	})
})
