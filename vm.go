package λ

import "errors"

type VM struct {
}

func (vm *VM) EvalString(src string) (error, Trace) {
	return errors.New("not implemented yet"), nil
}

var UnexpectedFreeVariable = errors.New("unexpected free variable")

type Trace interface {
	Pos() int
}
