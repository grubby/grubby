package testhelpers

import (
	"io/ioutil"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm"
	"github.com/grubby/grubby/interpreter/vm/builtins"
)

func SetupFileWithGlobalFilenameConst(vm vm.VM) {
	tempPath, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join(tempPath, "foo.rb"), []byte("foo = __FILE__"), 0600)

	if err != nil {
		panic(err)
	}

	loadPathGlobal, err := vm.Get("$LOAD_PATH")
	if err != nil {
		panic(err)
	}

	loadPathGlobal.(*builtins.Array).Append(builtins.NewString(tempPath))
}
