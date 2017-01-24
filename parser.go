package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, error) {
	var (
		scnr Scanner
		expr Λ
	)

	for {
		token := scnr.Scan(input)

		switch token {
		case LAMBDA, DOT, LPAREN, RPAREN:
		case EOF:
			return expr, nil
		default:
			thisVar := &Variable{string(token)}

			if expr == nil {
				expr = thisVar
			} else {
				expr = &Application{expr, thisVar}
			}
		}
	}
}

type layer struct {
}
