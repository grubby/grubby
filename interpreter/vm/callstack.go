package vm

import "fmt"

type CallStack struct {
	Frames []callStackFrame
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (stack *CallStack) Unshift(method, file string) {
	frame := callStackFrame{Method: method, File: file}
	stack.Frames = append([]callStackFrame{frame}, stack.Frames...)
}

func (stack *CallStack) Shift() {
	stack.Frames = stack.Frames[1:]
}

func (stack *CallStack) String() string {
	str := ""
	for _, frame := range stack.Frames {
		str += fmt.Sprintf("\t%s:in `%s'\n", frame.File, frame.Method)
	}

	return str
}

type callStackFrame struct {
	File   string
	Method string
}
