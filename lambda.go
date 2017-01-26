package λ

type Λ interface {
	Call(Λ) Λ
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
