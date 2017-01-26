package λ

type Λ interface {
	Call(Λ) Λ
}

func NewΛ(label string, fn func(μ Μ) Λ) Λ {
	return I
}

type Μ interface {
	NewΛ(string, func(μ Μ) Λ) Λ
	Read(string) Λ
}

type identity bool

func (_ *identity) Call(λ Λ) Λ {
	return λ
}

var I *identity = nil
