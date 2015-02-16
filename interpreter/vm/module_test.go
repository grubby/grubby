package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("modules", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be instantiated with the module keyword", func() {
		_, err := vm.Run(`
module Foo
end
`)

		Expect(err).ToNot(HaveOccurred())
		Expect(vm.Modules()["Foo"]).ToNot(BeNil())
	})

	It("has the Kernel#require method", func() {
		SetupLoadPathWithAFileToRequire(vm)
		_, err := vm.Run(`
module Foo
  require 'foo'
end
`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("has a .module_function method", func() {
		// this (historical) method can be used to mark instance methods as module
		// methods, so they can be included or extended elsewhere
		// (it incidentally should mark the original instance method as private)
		module, err := vm.Run(`
module Foo
  def whatever; end
  module_function :whatever

  alias something whatever
  module_function :something
end
`)

		Expect(err).ToNot(HaveOccurred())
		Expect(module).To(HaveMethod("whatever"))
	})

	It("supports the 'alias' keyword", func() {
		module, err := vm.Run(`
module Foo
  def from
  end

  alias to from
end
`)

		Expect(err).ToNot(HaveOccurred())
		Expect(module).To(HaveInstanceMethod("from"))
		Expect(module).To(HaveInstanceMethod("to"))
	})

	It("can be referred to inside of its declaration", func() {
		output := SwapStdout(func() {
			_, err := vm.Run(`
module Foo
  puts ::Foo
end
`)
			Expect(err).ToNot(HaveOccurred())
		})

		Expect(output).To(ContainSubstring("Foo"))
	})

	Describe("#module_eval", func() {
		It("can be used for evil", func() {
			value, err := vm.Run(`
class Foo
end

Foo.module_eval("def srs; 'bznz'; end")
Foo.new.srs
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(ContainSubstring("bznz"))
		})
	})
})
