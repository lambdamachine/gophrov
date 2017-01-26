package λ

import "fmt"

type Λ interface {
	String() string
}

type Variable struct {
	Name string
}

func (v *Variable) String() string {
	return v.Name
}

type Abstraction struct {
	Arg  *Variable
	Body Λ
}

func (λ *Abstraction) String() string {
	return fmt.Sprintf("λ%s.%s", λ.Arg, λ.Body)
}

type Application struct {
	Fn, Arg Λ
}

func (app *Application) String() string {
	return fmt.Sprintf("%s %s", &prettyFn{app.Fn}, &prettyArg{app.Arg})
}

type prettyFn struct {
	expr Λ
}

func (pty *prettyFn) String() string {
	switch pty.expr.(type) {
	case *Abstraction:
		return fmt.Sprintf("(%s)", pty.expr.String())
	default:
		return pty.expr.String()
	}
}

type prettyArg struct {
	expr Λ
}

func (pty *prettyArg) String() string {
	switch pty.expr.(type) {
	case *Variable:
		return pty.expr.String()
	default:
		return fmt.Sprintf("(%s)", pty.expr.String())
	}
}
