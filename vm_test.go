package λ_test

import (
	"."
	"testing"
)

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

func TestVM(t *testing.T) {
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
