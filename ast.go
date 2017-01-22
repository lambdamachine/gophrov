package λ

type Λ interface {
	String() string
}

type Variable struct {
	Name string
}

func (v *Variable) String() string {
	return ""
}

type Abstraction struct {
	Arg  *Variable
	Body Λ
}

func (λ *Abstraction) String() string {
	return ""
}

type Application struct {
	Fn, Arg Λ
}

func (app *Application) String() string {
	return ""
}
