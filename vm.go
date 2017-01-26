package λ

import (
	"bufio"
	"bytes"
	"errors"
)

type VM struct {
	Tape   Tape
	expr   Λ
	parser Parser
}

func (vm *VM) Quantum() Λ {
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

	expr = vm.reduce(expr)

	if nil == vm.expr {
		vm.expr = expr
	} else {
		vm.expr = &Application{vm.expr, expr}
	}

	vm.expr = vm.reduce(vm.expr)

	return nil, nil
}

type Tape interface {
	Load(string) Λ
	Expand(string, Λ)
	Shrink()
}

type tape struct {
	data map[string]Λ
}

func newTape() *tape {
	return &tape{data: map[string]Λ{}}
}

func (t *tape) Load(key string) (expr Λ) {
	expr, _ = t.data[key]
	return
}

func (t *tape) Expand(key string, expr Λ) {
	t.data[key] = expr
}

func (t *tape) Shrink() {

}

func (vm *VM) reduce(expr Λ) Λ {
	if nil == vm.Tape {
		vm.Tape = newTape()
	}

	stack := []Λ{}

	for {
		switch current := expr.(type) {
		case *Abstraction:
			vm.Tape.Shrink()

			if len(stack) > 0 {
				expr, stack = stack[len(stack)-1], stack[:len(stack)-1]
				expr = &Application{current, expr}
				continue
			}

			return expr
		case *Application:
			switch fn := current.Fn.(type) {
			case *Abstraction:
				switch arg := current.Arg.(type) {
				case *Variable:
					vm.Tape.Expand(fn.Arg.Name, vm.Tape.Load(arg.Name))
				default:
					vm.Tape.Expand(fn.Arg.Name, current.Arg)
				}

				expr = fn.Body
			case *Application:
				stack = append(stack, current.Arg)
				expr = fn
			case *Variable:
				current.Fn = vm.Tape.Load(fn.Name)
			}
		case *Variable:
			expr = vm.Tape.Load(current.Name)
		}
	}
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
