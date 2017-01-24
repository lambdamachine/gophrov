package λ_test

import (
	"."
	"bufio"
	"bytes"
	"testing"
)

var parserValidExamples = map[string]string{
	"x":                     "x",
	"x y":                   "x y",
	"x y z":                 "x y z",
	"(x)":                   "x",
	"(x y)":                 "x y",
	"(x y z)":               "x y z",
	"((x y) z)":             "x y z",
	"x (y z)":               "x (y z)",
	"((x (y z) (x y z) y))": "x (y z) (x y z) y",
	"x ((y) (((z x))))":     "x (y (z x))",
	"λx.x":                  "λx.x",
	"λx.λy.x":               "λx.λy.x",
	"(λx.x)":                "λx.x",
	"(λx.(λy.x))":           "λx.λy.x",
	"λx.x (y z) m n":        "λx.x (y z) m n",
	// FIXME: "x (λx.x y (λy.x) z)":   "x (λx.x y (λy.x) z)",
}

func TestParseValidLambdaExpressions(t *testing.T) {
	var parser λ.Parser

	for example, expectedExprString := range parserValidExamples {
		source := bufio.NewReader(bytes.NewReader([]byte(example)))
		expr, n, err := parser.Parse(source)

		if err != nil {
			t.Errorf("expected (%s) to be parsed successfully, got error: %v",
				example, err)
		} else if expr.String() != expectedExprString || n != len(example) {
			t.Errorf("expected to parse %d runes of (%s) as (%s), got %d runes as (%s) instead",
				len(example), example, expectedExprString, n, expr)
		}
	}
}

var parserInvalidExamples = map[string]struct {
	err error
	pos int
}{
	"":           {λ.UnexpectedEndOfInput, 0},
	"(":          {λ.UnexpectedEndOfInput, 1},
	"(x":         {λ.UnexpectedEndOfInput, 2},
	"(x y (x y)": {λ.UnexpectedEndOfInput, 10},
	")":          {λ.UnexpectedToken, 0},
	"()":         {λ.UnexpectedToken, 1},
	"λ)":         {λ.UnexpectedToken, 2},
}

func TestParseInvalidLambdaExpressions(t *testing.T) {
	var parser λ.Parser

	for example, expected := range parserInvalidExamples {
		source := bufio.NewReader(bytes.NewReader([]byte(example)))
		_, pos, err := parser.Parse(source)

		if err == nil {
			t.Errorf("expression (%s) should be recognized as invalid ",
				example)
		} else if err.Error() != expected.err.Error() || pos != expected.pos {
			t.Errorf("expected (%s) to throw %#v at rune %d, got %#v at rune %d instead",
				example, expected.err, expected.pos, err, pos)
		}
	}
}
