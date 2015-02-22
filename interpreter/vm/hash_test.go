package vm_test

import (
	"os"
	"path/filepath"

	. "github.com/grubby/grubby/interpreter/vm"
	. "github.com/grubby/grubby/interpreter/vm/builtins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hashes", func() {
	var vm VM

	BeforeEach(func() {
		pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))))
		if err != nil {
			panic(err)
		}

		vm = NewVM(pathToExecutable, "fake-irb-under-test")
	})

	It("can be constructed as a hash literal", func() {
		hash, err := vm.Run("{:key => :value}")
		Expect(err).ToNot(HaveOccurred())

		keysMethod, err := hash.Method("keys")
		Expect(err).ToNot(HaveOccurred())

		keys, err := keysMethod.Execute(hash, nil)
		Expect(err).ToNot(HaveOccurred())

		keyArray, ok := keys.(*Array)
		Expect(ok).To(BeTrue())
		Expect(keyArray.Members()).To(ContainElement(vm.Symbols()["key"]))

		valuesMethod, err := hash.Method("values")
		Expect(err).ToNot(HaveOccurred())

		values, err := valuesMethod.Execute(hash, nil)
		Expect(err).ToNot(HaveOccurred())

		valueArray, ok := values.(*Array)
		Expect(ok).To(BeTrue())
		Expect(valueArray.Members()).To(ContainElement(vm.Symbols()["value"]))

		method, err := hash.Method("[]=")
		Expect(err).ToNot(HaveOccurred())

		fooSymbol := NewSymbol("foo", vm)
		barSymbol := NewSymbol("bar", vm)

		_, err = method.Execute(hash, nil, fooSymbol, barSymbol)
		Expect(err).ToNot(HaveOccurred())

		keys, err = keysMethod.Execute(hash, nil)
		Expect(err).ToNot(HaveOccurred())

		keyArray, ok = keys.(*Array)
		Expect(ok).To(BeTrue())
		Expect(keyArray.Members()).To(ContainElement(fooSymbol))

		values, err = valuesMethod.Execute(hash, nil)
		Expect(err).ToNot(HaveOccurred())

		valueArray, ok = values.(*Array)
		Expect(ok).To(BeTrue())
		Expect(valueArray.Members()).To(ContainElement(barSymbol))
	})

	It("has a reasonable String() method", func() {
		val, err := vm.Run("{:hello => 'world'}")
		Expect(err).ToNot(HaveOccurred())
		Expect(val.String()).To(Equal("{:hello => \"world\"}"))
	})

	It("can yield the value for keys with the [] operator", func() {
		value, err := vm.Run(`
hash = {:hello => :world}
hash[:hello]
`)
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(vm.Symbols()["world"]))
	})

	Describe("iterating over the keys and items", func() {
		var err error

		BeforeEach(func() {
			_, err = vm.Run(`
keys = []
values = []

{:foo => 1, :bar => 2}.each do |key, value|
  keys.unshift(key)
  values.unshift(value)
end
`)
		})

		It("calls the provided block once with each item in the hash", func() {
			Expect(err).ToNot(HaveOccurred())

			keysObj := vm.MustGet("keys")
			keysArray, ok := keysObj.(*Array)
			Expect(ok).To(BeTrue())

			keys := keysArray.Members()

			Expect(len(keys)).To(Equal(2))
			Expect(keys).To(ContainElement(vm.Symbols()["foo"]))
			Expect(keys).To(ContainElement(vm.Symbols()["bar"]))

			valuesObj := vm.MustGet("values")
			valuesArray, ok := valuesObj.(*Array)
			Expect(ok).To(BeTrue())

			values := valuesArray.Members()

			Expect(len(values)).To(Equal(2))
			Expect(values).To(ContainElement(NewFixnum(1, vm, vm)))
			Expect(values).To(ContainElement(NewFixnum(2, vm, vm)))
		})
	})
})
