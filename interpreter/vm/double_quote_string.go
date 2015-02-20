package vm

import (
	"strings"

	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretDoubleQuotedStringInContext(
	vm *vm,
	stringValue ast.InterpolatedString,
	context Value,
) (Value, error) {

	str := stringValue.Value
	currentBytes := []byte{}
	bytesToInterpret := [][]byte{}
	insideInterpolation := false
	for i := 0; i < len(str); i++ {
		if insideInterpolation {
			currentBytes = append(currentBytes, str[i])
		}

		if str[i] == '}' && i > 0 && str[i-1] != '\\' {
			insideInterpolation = false
			bytesToInterpret = append(bytesToInterpret, currentBytes)
			currentBytes = []byte{}
		} else if str[i] == '#' && len(str) > i && str[i+1] == '{' {
			insideInterpolation = true
		}
	}

	for _, bytes := range bytesToInterpret {
		substringToReplace := string(bytes)
		rubyValue, err := vm.EvaluateStringInContext(substringToReplace[1:len(substringToReplace)-1], context)
		if err != nil {
			return nil, err
		}

		str = strings.Replace(str, "#"+substringToReplace, rubyValue.String(), 1)
	}

	return NewString(str, vm, vm), nil
}
