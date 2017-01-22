package λ_test

import (
	"."
	"testing"
)

var astExamples = map[λ.Λ]string{
	&λ.Abstraction{
		&λ.Variable{"x"},
		&λ.Variable{"x"},
	}: "λx.x",
	&λ.Abstraction{
		&λ.Variable{"x"},
		&λ.Application{
			&λ.Variable{"x"},
			&λ.Variable{"x"},
		},
	}: "λx.x x",
	&λ.Abstraction{
		&λ.Variable{"hello"},
		&λ.Application{
			&λ.Application{
				&λ.Variable{"hello"},
				&λ.Variable{"there"},
			},
			&λ.Variable{"world"},
		},
	}: "λhello.hello there world",
}

func TestAST(t *testing.T) {
	for expr, expected := range astExamples {
		if expr.String() != expected {
			t.Errorf("expected (%v), got (%s)", expected, expr)
		}
	}
}
