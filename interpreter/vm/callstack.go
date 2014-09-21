package vm

import "fmt"

type CallStack struct {
	Lines []string
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (stack *CallStack) Unshift(line string) {
	stack.Lines = append([]string{line}, stack.Lines...)
}

func (stack *CallStack) Shift() {
	stack.Lines = stack.Lines[1:]
}

func (stack *CallStack) String() string {
	str := ""
	for _, line := range stack.Lines {
		str += fmt.Sprintf("\t%s\n", line)
	}

	return str
}
