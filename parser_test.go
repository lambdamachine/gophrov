package λ_test

import (
	"."
	"bufio"
	"bytes"
	"testing"
)

var parserValidExamples = map[string]string{
	"x": "x",
}

func TestParseValidLambdaExpressions(t *testing.T) {
	var parser λ.Parser

	for example, expectedExprString := range parserValidExamples {
		source := bufio.NewReader(bytes.NewReader([]byte(example)))
		expr, err := parser.Parse(source)

		if err != nil {
			t.Errorf("expected (%s) to be parsed successfully, got error: %v",
				example, err)
		} else if expr.String() != expectedExprString {
			t.Errorf("expected (%s) to be parsed as (%s), got (%s) instead",
				example, expectedExprString, expr)
		}
	}
}
