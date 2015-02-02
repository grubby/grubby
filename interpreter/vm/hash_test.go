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
		val, err := vm.Run("{:hello => 'world', :foo => :bar}")
		Expect(err).ToNot(HaveOccurred())
		Expect(val.String()).To(Equal("{:hello => \"world\", :foo => :bar}"))
	})
})
