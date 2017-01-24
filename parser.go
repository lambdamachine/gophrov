package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, error) {
	var (
		scnr Scanner
		zn   = &zone{nil, nil}
	)

	for {
		token := scnr.Scan(input)

		switch token {
		case LPAREN:
			zn = &zone{zn, nil}
		case RPAREN:
			expr := zn.expr
			zn = zn.zn

			if zn.expr != nil {
				expr = &Application{zn.expr, expr}
			}

			zn.expr = expr
		case LAMBDA:
			token = scnr.Scan(input)

			switch token {
			case LAMBDA, LPAREN, RPAREN, DOT, EOF:
			default:
				zn.expr = &Abstraction{&Variable{string(token)}, nil}
				zn = &zone{zn, nil}
			}
		case DOT:
		case EOF:
			expr := zn.expr

			for zn.zn != nil {
				if abstr, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
					zn = zn.zn
					abstr.Body = expr
					expr = zn.expr
				}
			}

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
	zn   *zone
	expr Λ
}
