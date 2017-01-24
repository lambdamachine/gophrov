package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, int, error) {
	var (
		scnr Scanner
		pos  = 0
		zn   = &zone{nil, nil}
	)

	for {
		token, n := scnr.Scan(input)
		pos += n

		switch token {
		case LPAREN:
			zn = &zone{zn, nil}
		case RPAREN:
			if zn.expr == nil {
				return nil, pos - 1, UnexpectedToken
			}

			expr := zn.expr

			for zn.zn != nil {
				if abstr, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
					zn = zn.zn
					abstr.Body = expr
					expr = zn.expr
				} else {
					break
				}
			}

			zn = zn.zn

			if zn.expr != nil {
				expr = &Application{zn.expr, expr}
			}

			zn.expr = expr
		case LAMBDA:
			token, n = scnr.Scan(input)
			pos += n

			switch token {
			case LAMBDA, LPAREN, RPAREN, DOT, EOF:
				return nil, pos - 1, UnexpectedToken
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
				} else {
					break
				}
			}

			if zn.expr == nil || zn.zn != nil {
				return nil, pos, UnexpectedEndOfInput
			}

			return zn.expr, pos, nil
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

type unexpectedEndOfInput struct{}

func (err *unexpectedEndOfInput) Error() string {
	return "unexpected end of input"
}

func (err *unexpectedEndOfInput) GoString() string {
	return "UnexpectedEndOfInput"
}

var UnexpectedEndOfInput = &unexpectedEndOfInput{}

type unexpectedToken struct{}

func (err *unexpectedToken) Error() string {
	return "unexpected token"
}

func (err *unexpectedToken) GoString() string {
	return "UnexpectedToken"
}

var UnexpectedToken = &unexpectedToken{}
