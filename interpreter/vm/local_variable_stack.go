package vm

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/interpreter/vm/builtins"
)

type frame map[string]builtins.Value

type LocalVariableStack struct {
	frames []frame
}

func NewLocalVariableStack() *LocalVariableStack {
	return &LocalVariableStack{
		frames: make([]frame, 0),
	}
}

func newEmptyFrame() []frame {
	return []frame{{}}
}

func (stack *LocalVariableStack) Unshift() {
	stack.frames = append(newEmptyFrame(), stack.frames...)
}

func (stack *LocalVariableStack) Shift() {
	stack.frames = stack.frames[0 : len(stack.frames)-1]
}

func (stack *LocalVariableStack) UnshiftCopyingCurrentFrame() {
	var lastFrame frame
	if len(stack.frames) >= 1 {
		lastFrame = stack.frames[0]
	} else {
		lastFrame = newEmptyFrame()[0]
	}

	newFrame := newEmptyFrame()
	for key, value := range lastFrame {
		newFrame[0][key] = value
	}

	stack.frames = append(newFrame, stack.frames...)
}

func (stack *LocalVariableStack) Store(key string, value builtins.Value) {
	stack.frames[0][key] = value
}

func (stack *LocalVariableStack) Retrieve(key string) (builtins.Value, error) {
	maybe, ok := stack.frames[0][key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("No such key '%s'", key))
	} else {
		return maybe, nil
	}
}
