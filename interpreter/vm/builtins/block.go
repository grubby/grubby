package builtins

import "github.com/grubby/grubby/ast"

type BlockEvaluator interface {
	EvaluateBlockWithArgsInContext(Value, []BlockArg, []ast.Node) (Value, error)
}

type Block interface {
	Call(args ...Value) (Value, error)
}

type BlockArg struct {
	Name  string
	Value Value
}

type blockImpl struct {
	valueStub

	Context   Value
	args      []ast.Node
	body      []ast.Node
	evaluator BlockEvaluator
}

func (b *blockImpl) Call(args ...Value) (Value, error) {
	invocationArgs := make([]BlockArg, 0, len(args))
	for index, providedArg := range args {
		blockArg := BlockArg{
			Name:  b.args[index].(ast.BareReference).Name,
			Value: providedArg,
		}
		invocationArgs = append(invocationArgs, blockArg)
	}

	return b.evaluator.EvaluateBlockWithArgsInContext(b.Context, invocationArgs, b.body)
}

func NewBlock(Context Value, args []ast.Node, body []ast.Node, evaluator BlockEvaluator) Block {
	return &blockImpl{
		Context:   Context,
		args:      args,
		body:      body,
		evaluator: evaluator,
	}
}
