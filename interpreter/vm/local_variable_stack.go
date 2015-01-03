package vm

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/interpreter/vm/builtins"
)

type frame map[string]builtins.Value

type localVariableStack struct {
	frames []frame
}

func newLocalVariableStack() *localVariableStack {
	return &localVariableStack{
		frames: make([]frame, 0),
	}
}

func newEmptyFrame() []frame {
	return []frame{
		frame{},
	}
}

func (stack *localVariableStack) unshift() {
	stack.frames = append(newEmptyFrame(), stack.frames...)
}

func (stack *localVariableStack) shift() {
	stack.frames = stack.frames[0 : len(stack.frames)-1]
}

func (stack *localVariableStack) store(key string, value builtins.Value) {
	stack.frames[0][key] = value
}

func (stack *localVariableStack) retrieve(key string) (builtins.Value, error) {
	for _, frame := range stack.frames {
		val, ok := frame[key]
		if ok {
			return val, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No such key '%s'", key))
}
