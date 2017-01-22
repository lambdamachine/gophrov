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
	return fmt.Sprintf("%s %s", app.Fn, app.Arg)
}
