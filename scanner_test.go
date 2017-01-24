package λ_test

import (
	"."
	"bufio"
	"bytes"
	"testing"
)

var scannerExamples = map[string][]λ.Token{
	"": {
		λ.EOF,
	},
	"   ": {
		λ.EOF,
	},
	" \n\n \t ": {
		λ.EOF,
	},
	"λ.": {
		λ.LAMBDA,
		λ.DOT,
		λ.EOF,
	},
	"λ\n. ": {
		λ.LAMBDA,
		λ.DOT,
		λ.EOF,
	},
	" λ   \n. \t\n": {
		λ.LAMBDA,
		λ.DOT,
		λ.EOF,
	},
	"()": {
		λ.LPAREN,
		λ.RPAREN,
		λ.EOF,
	},
	"  (  )  ": {
		λ.LPAREN,
		λ.RPAREN,
		λ.EOF,
	},
	" \n\n ( \n ) \t\n\n": {
		λ.LPAREN,
		λ.RPAREN,
		λ.EOF,
	},
	"(λ.)": {
		λ.LPAREN,
		λ.LAMBDA,
		λ.DOT,
		λ.RPAREN,
		λ.EOF,
	},
	"(λx.x)": {
		λ.LPAREN,
		λ.LAMBDA,
		λ.Token("x"),
		λ.DOT,
		λ.Token("x"),
		λ.RPAREN,
		λ.EOF,
	},
	"λx.x": {
		λ.LAMBDA,
		λ.Token("x"),
		λ.DOT,
		λ.Token("x"),
		λ.EOF,
	},
	"λx.x (hello (yada yada)\n\tλyyy.yyy zzz)": {
		λ.LAMBDA,
		λ.Token("x"),
		λ.DOT,
		λ.Token("x"),
		λ.LPAREN,
		λ.Token("hello"),
		λ.LPAREN,
		λ.Token("yada"),
		λ.Token("yada"),
		λ.RPAREN,
		λ.LAMBDA,
		λ.Token("yyy"),
		λ.DOT,
		λ.Token("yyy"),
		λ.Token("zzz"),
		λ.RPAREN,
		λ.EOF,
	},
}

func TestScanner(t *testing.T) {
	var scanner λ.Scanner

	for example, expected := range scannerExamples {
		input := bufio.NewReader(bytes.NewReader([]byte(example)))

		for i := 0; ; i++ {
			token := scanner.Scan(input)
			expectedToken := expected[i]

			if token != expectedToken {
				t.Fatalf("expected to scan '%s', got '%s'", expectedToken, token)
			}

			if token == λ.EOF {
				if len(expected) > i+1 {
					t.Fatalf("not enough tokens has been scanned")
				}

				break
			}
		}
	}
}
