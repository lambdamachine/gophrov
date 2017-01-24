package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, error) {
	var (
		scnr Scanner
		zn   = &zone{nil}
	)

	for {
		token := scnr.Scan(input)

		switch token {
		case LAMBDA, DOT, LPAREN, RPAREN:
		case EOF:
			return zn.expr, nil
		default:
			thisVar := &Variable{string(token)}

			if zn.expr == nil {
				zn.expr = thisVar
			} else {
				zn.expr = &Application{zn.expr, thisVar}
			}
		}
	}
}

type zone struct {
  expr Λ
}
