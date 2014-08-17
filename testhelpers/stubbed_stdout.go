package testhelpers

import (
	"bufio"
	"os"

	. "github.com/onsi/gomega"
)

func SwapStdout(block func()) string {
	in, out, err := os.Pipe()
	Expect(err).ToNot(HaveOccurred())

	oldPipe := os.Stdout
	defer func() {
		os.Stdout = oldPipe
	}()
	os.Stdout = out

	block()

	reader := bufio.NewReader(in)
	str, err := reader.ReadString('\n')

	Expect(err).ToNot(HaveOccurred())
	return str
}
