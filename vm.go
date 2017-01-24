package λ

import (
	"bufio"
	"bytes"
	"errors"
)

type VM struct {
	expr   Λ
	parser Parser
}

func (vm *VM) EvalString(src string) (error, Trace) {
	vm.parser.Report = func(r Report) error {
		return nil
	}

	input := bufio.NewReader(bytes.NewReader([]byte(src)))
	expr, pos, err := vm.parser.Parse(input)

	if err != nil {
		return err, &trace{pos: pos}
	}

	vm.expr = expr

	return nil, nil
}

var UnexpectedFreeVariable = errors.New("unexpected free variable")

type Trace interface {
	Pos() int
}

type trace struct {
	pos int
}

func (trc *trace) Pos() int {
	return trc.pos
}
