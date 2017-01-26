package λ

import "fmt"

type Expression interface {
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
	Body Expression
}

func (λ *Abstraction) String() string {
	return fmt.Sprintf("λ%s.%s", λ.Arg, λ.Body)
}

type Application struct {
	Fn, Arg Expression
}

func (app *Application) String() string {
	return fmt.Sprintf("%s %s", app.Fn, &parens{app.Arg})
}

type parens struct {
	expr Expression
}

func (prns *parens) String() string {
	switch prns.expr.(type) {
	case *Variable:
		return prns.expr.String()
	default:
		return fmt.Sprintf("(%s)", prns.expr.String())
	}
}
