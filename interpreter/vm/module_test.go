package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
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

  private
  def private_whatever; end
  module_function :private_whatever
end
`)

		Expect(err).ToNot(HaveOccurred())
		Expect(module).To(HaveMethod("whatever"))
		Expect(module).To(HaveMethod("private_whatever"))
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

	It("supports aliasing methods with #alias_method", func() {
		module, err := vm.Run(`
module Foo
  def foo; end
  alias_method :bar, :foo
end
`)

		Expect(err).ToNot(HaveOccurred())
		Expect(module).To(HaveInstanceMethod("foo"))
		Expect(module).To(HaveInstanceMethod("bar"))
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

Foo.module_eval("def srs()\n'bznz'\nend")
Foo.new.srs
`)

			Expect(err).ToNot(HaveOccurred())
			Expect(value.String()).To(ContainSubstring("bznz"))
		})
	})

	Describe("nested modules", func() {
		It("can be used to refer to constants defined in the parent module", func() {
			value, err := vm.Run(`
module Foo
  FOO_CONSTANT = 'hello'
  module Bar
    def self.test
      ::Foo::FOO_CONSTANT
    end
  end
end

::Foo::Bar.test
`)
			Expect(err).ToNot(HaveOccurred())
			Expect(value.(*StringValue).RawString()).To(Equal("hello"))
		})

		It("can be referred to by the parent module", func() {
			value, err := vm.Run(`
module Foo
  module Bar
    MESSAGE = 'drink your ovaltine'
  end

  include Bar
end
`)

			Expect(err).ToNot(HaveOccurred())

			module, ok := value.(Module)
			Expect(ok).To(BeTrue())

			constant, err := module.Constant("MESSAGE")
			Expect(err).ToNot(HaveOccurred())

			stringValue, ok := constant.(*StringValue)
			Expect(ok).To(BeTrue())
			Expect(stringValue.RawString()).To(Equal("drink your ovaltine"))
		})
	})

	It("can be referred to by modules declared alongside it", func() {
		value, err := vm.Run(`
module Foo
  module Bar
    MESSAGE = 'drink your ovaltine'
  end

  module Baz
    include Bar
  end
end

Foo::Baz::MESSAGE
`)

		Expect(err).ToNot(HaveOccurred())

		stringValue, ok := value.(*StringValue)
		Expect(ok).To(BeTrue())
		Expect(stringValue.RawString()).To(Equal("drink your ovaltine"))
	})
})
