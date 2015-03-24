package vm

import (
	"strings"

	"github.com/grubby/grubby/ast"
	"github.com/tjarratt/gomads"

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

		valueAsString := gomads.Maybe(func() interface{} {
			method, err := rubyValue.Method("to_s")
			if err != nil {
				return nil
			}

			result, err := method.Execute(rubyValue, nil)
			if err != nil {
				return nil
			}

			return result.(*StringValue).RawString()
		}).OrSome(rubyValue.String()).Value().(string)

		str = strings.Replace(str, "#"+substringToReplace, valueAsString, 1)
	}

	return NewString(str, vm, vm), nil
}
