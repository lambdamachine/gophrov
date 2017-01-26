package λ

import "errors"

type Λ interface {
	Call(Λ) Λ
}

func NewΛ(label string, fn func(μ Μ) Λ) Λ {
	return &combinator{nil, &meta{label, fn}, nil}
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

type combinator struct {
	parent *combinator
	meta   *meta
	λ      Λ
}

type meta struct {
	label string
	fn    func(μ Μ) Λ
}

func (cbntr *combinator) NewΛ(label string, fn func(μ Μ) Λ) Λ {
	return &combinator{cbntr, &meta{label, fn}, nil}
}

func (cbntr *combinator) Call(λ Λ) Λ {
	return cbntr.meta.fn(&combinator{cbntr, cbntr.meta, λ})
}

func (cbntr *combinator) Read(label string) Λ {
	for current := cbntr; current != nil; current = current.parent {
		if current.meta.label == label {
			return current.λ
		}
	}

	panic(InconsistentMemory)
}

var InconsistentMemory = errors.New("inconsistent memory")
