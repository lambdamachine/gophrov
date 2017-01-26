package λ_test

import (
	"."
	"testing"
)

var vmValidExamples = []struct {
	source  string
	quantum string
}{
	{
		"λx.x",
		"λx.x",
	},
	{
		"λx.λy.x",
		"λx.λy.x",
	},
	{
		"λx.x",
		"λy.x",
	},
	{
		"λx y z.x",
		"λx.x",
	},
	{
		"λx y.y",
		"λx.λy.y",
	},
	{
		"λx y.x y",
		"λy.y",
	},
	{
		"λx.λx.x",
		"λx.λx.x",
	},
	{
		"λa.a",
		"λx.x",
	},
	// {
	// 	"(λx y.x) (λc.c) (λa b.b)",
	// 	"λc.c",
	// },
	{
		"(λa b.a (λa.a) a) (λz.z) (λa.a)",
		"λz.z",
	},
	{
		"(λa.((λa.a) (λf.f)) ((λx.x) a)) (λt.t)",
		"λt.t",
	},
	{
		"λx.λy.λz.x y (y z)",
		"λx.λy.λz.x y (y z)",
	},
}

func TestVMEvaluationSuccesses(t *testing.T) {
	var vm λ.VM

	for _, example := range vmValidExamples {
		err, _ := vm.EvalString(example.source)

		if err != nil {
			t.Fatalf("expected successful evaluation of (%s), got error %q",
				example.source, err)
		} else if example.quantum != vm.Quantum().String() {
			t.Fatalf("expected quantum (%s), got (%s)", example.quantum, vm.Quantum())
		}
	}
}

var vmInvalidExamples = map[string]struct {
	err error
	pos int
}{
	"x":             {λ.UnexpectedFreeVariable, 0},
	"λx.y":          {λ.UnexpectedFreeVariable, 3},
	"λx y.x y z":    {λ.UnexpectedFreeVariable, 8},
	"λx.λy.x y z":   {λ.UnexpectedFreeVariable, 9},
	"λx.x (x λy.z)": {λ.UnexpectedFreeVariable, 11},
	"λx..":          {λ.UnexpectedToken, 3},
}

func TestVMEvaluationErrors(t *testing.T) {
	var vm λ.VM

	for example, expected := range vmInvalidExamples {
		err, trace := vm.EvalString(example)

		if err == nil {
			t.Errorf("expression (%s) should not be evaluated", example)
		} else if trace == nil {
			t.Errorf("error trace is missing")
		} else if err != expected.err || trace.Pos() != expected.pos {
			t.Errorf("eval of (%s) was expected to throw %q at rune %d, got %q at rune %d instead",
				example, expected.err, expected.pos, err, trace.Pos())
		}
	}
}
