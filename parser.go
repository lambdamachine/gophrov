package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, int, error) {
	var (
		scnr Scanner
		pos  = 0
		zn   = &zone{nil, 0, nil}
	)

	for {
		token, n := scnr.Scan(input)
		pos += n

		switch token {
		case LPAREN:
			zn = &zone{zn, zn.paren + 1, nil}
		case RPAREN:
			if zn.expr == nil {
				return nil, pos - 1, UnexpectedToken
			}

			for paren := zn.paren; zn.paren == paren; {
				var expr Λ
				zn, expr = closeAbstractions(zn)

				if zn.expr == nil || zn.zn == nil {
					return nil, pos - 1, UnexpectedToken
				}

				zn = zn.zn

				if zn.expr != nil {
					expr = &Application{zn.expr, expr}
				}

				zn.expr = expr
			}
		case LAMBDA:
			if zn.expr != nil {
				zn = &zone{zn, zn.paren, nil}
			}

		definition:
			for {
				token, n = scnr.Scan(input)
				pos += n

				switch token {
				case LAMBDA, LPAREN, RPAREN:
					return nil, pos - 1, UnexpectedToken
				case DOT:
					if zn.zn != nil {
						if _, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
							break definition
						}
					}

					return nil, pos - 1, UnexpectedToken
				case EOF:
					return nil, pos, UnexpectedEndOfInput
				default:
					zn.expr = &Abstraction{&Variable{string(token)}, nil}
					zn = &zone{zn, zn.paren, nil}
				}
			}
		case DOT:
			return nil, pos - 1, UnexpectedToken
		case EOF:
			if zn.expr == nil {
				return nil, pos, UnexpectedEndOfInput
			}

			zn, _ = closeAbstractions(zn)

			if zn.expr == nil || zn.zn != nil {
				return nil, pos, UnexpectedEndOfInput
			}

			return zn.expr, pos, nil
		default:
			var expr Λ = &Variable{string(token)}

			if zn.expr != nil {
				expr = &Application{zn.expr, expr}
			}

			zn.expr = expr
		}
	}
}

func closeAbstractions(inzn *zone) (zn *zone, expr Λ) {
	zn, expr = inzn, inzn.expr

	for zn.zn != nil {
		if abstr, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
			zn = zn.zn
			abstr.Body = expr
			expr = zn.expr
		} else {
			break
		}
	}

	return
}

type zone struct {
	zn    *zone
	paren int
	expr  Λ
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
