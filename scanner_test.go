package λ_test

import (
	"."
	"bytes"
	"testing"
)

var scannerExamples = map[string][]λ.Token{
	"": {
		λ.EOF,
	},
}

func TestScanner(t *testing.T) {
	for example, expected := range scannerExamples {
		input := bytes.NewReader([]byte(example))
		scanner := λ.NewScanner(input)

		for i := 0; ; i++ {
			token := scanner.Scan()
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
