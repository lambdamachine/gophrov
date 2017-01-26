package λ

import (
	"bufio"
	"bytes"
	"errors"
	"time"
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
	Save(string, Λ)
	Expand(int)
	Shrink(int)
}

type track struct {
	parent *track
	lvl    int
	m      map[string]Λ
}

func newTrack(parent *track, lvl int) *track {
	return &track{parent, lvl, map[string]Λ{}}
}

type tape struct {
	track *track
}

func newTape() *tape {
	return &tape{track: newTrack(nil, 0)}
}

func (t *tape) Load(key string) (expr Λ) {
	var ok bool

	for head := t.track; head != nil; head = head.parent {
		if expr, ok = head.m[key]; ok {
			return
		}
	}

	return
}

func (t *tape) Save(key string, expr Λ) {
	t.track.m[key] = expr
}

func (t *tape) Expand(lvl int) {
	t.track = newTrack(t.track, lvl)
}

func (t *tape) Shrink(lvl int) {
	if nil != t.track {
		t.track = t.track.parent
	}
}

func (vm *VM) reduce(expr Λ) Λ {
	if nil == vm.Tape {
		vm.Tape = newTape()
	}

	stack := []Λ{}

	for {
		switch current := expr.(type) {
		case *Abstraction:
			if len(stack) > 0 {
				switch current.Body.(type) {
				case *Abstraction:
					vm.Tape.Shrink(0)
				}

				expr, stack = stack[len(stack)-1], stack[:len(stack)-1]
				expr = &Application{current, expr}
				continue
			}

			return expr
		case *Application:
			switch fn := current.Fn.(type) {
			case *Abstraction:
				vm.Tape.Expand(len(stack))

				switch arg := current.Arg.(type) {
				case *Variable:
					vm.Tape.Save(fn.Arg.Name, vm.Tape.Load(arg.Name))
				default:
					vm.Tape.Save(fn.Arg.Name, current.Arg)
				}

				expr = fn.Body
			case *Application:
				stack = append(stack, current.Arg)
				expr = fn
			case *Variable:
				expr = &Application{vm.Tape.Load(fn.Name), current.Arg}
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
