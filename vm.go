package λ

import (
	"bufio"
	"bytes"
	"errors"
)

type VM struct {
	expr   Expression
	parser Parser
}

func (vm *VM) Quantum() Expression {
	return vm.expr
}

func (vm *VM) EvalString(src string) (error, Trace) {
	scp := &scope{nil, map[string]bool{}}

	vm.parser.Report = func(r Report) (err error) {
		if r == nil {
			return
		}

		switch r.Event() {
		case ABSTRACTION_ENTER:
			switch expr := r.Expr().(type) {
			case *Abstraction:
				scp = scp.NewNestedScope(expr.Arg.Name)
			}
		case ABSTRACTION_EXIT:
			scp = scp.parent
		case VARIABLE_SPOT:
			switch expr := r.Expr().(type) {
			case *Variable:
				if !scp.HasName(expr.Name) {
					return UnexpectedFreeVariable
				}
			}
		}

		return
	}

	input := bufio.NewReader(bytes.NewReader([]byte(src)))
	expr, pos, err := vm.parser.Parse(input)

	if err != nil {
		return err, &trace{pos: pos}
	}

	if nil == vm.expr {
		vm.expr = expr
	} else {
		vm.expr = &Application{vm.expr, expr}
	}

	vm.reduce()

	return nil, nil
}

func (vm *VM) reduce() {
	// ???
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

type scope struct {
	parent *scope
	names  map[string]bool
}

func (scp *scope) NewNestedScope(name string) *scope {
	names := map[string]bool{}

	for k, v := range scp.names {
		names[k] = v
	}

	names[name] = true
	return &scope{scp, names}
}

func (scp *scope) HasName(name string) (ok bool) {
	_, ok = scp.names[name]
	return
}
