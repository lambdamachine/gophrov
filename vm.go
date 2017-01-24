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
	sv := make(chan Report)
	vm.parser.Supervisor = sv

	defer close(sv)

	go func() {
		for r := range sv {
			if r == nil {
				break
			}

			switch r.Event() {
			case ABSTRACTION_ENTER:

			case ABSTRACTION_EXIT:

			case VARIABLE_SPOT:

			}
		}
	}()

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
