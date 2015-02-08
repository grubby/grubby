package vm

import "fmt"

type CallStack struct {
	Frames []callStackFrame
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (stack *CallStack) Unshift(method, file string, lineNumber int) {
	frame := callStackFrame{Method: method, File: file, LineNumber: lineNumber}
	stack.Frames = append([]callStackFrame{frame}, stack.Frames...)
}

func (stack *CallStack) Shift() {
	stack.Frames = stack.Frames[1:]
}

func (stack *CallStack) String() string {
	str := ""
	for _, frame := range stack.Frames {
		str += fmt.Sprintf("\t%s:%d in `%s'\n", frame.File, frame.LineNumber+1, frame.Method)
	}

	return str
}

type callStackFrame struct {
	File       string
	Method     string
	LineNumber int
}
