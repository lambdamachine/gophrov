package λ

import (
	"io"
)

type Parser struct {
	scnr Scanner
}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, int, error) {
	var (
		pos = 0
		zn  = newRootZone()
	)

	for {
		token, n := prsr.scnr.Scan(input)
		pos += n

		switch token {
		case LPAREN:
			zn = zn.NewParenZone()
		case RPAREN:
			if zn.IsEmpty() {
				return nil, pos - 1, UnexpectedToken
			}

			for paren := zn.paren; zn.paren == paren; {
				var expr Λ
				zn, expr = closeAbstractions(zn)

				if zn.IsEmpty() || zn.IsRoot() {
					return nil, pos - 1, UnexpectedToken
				}

				zn = zn.zn
				zn.SetOrApply(expr)
			}
		case LAMBDA:
			if !zn.IsEmpty() {
				zn = zn.NewAbstractionZone()
			}

		definition:
			for {
				token, n = prsr.scnr.Scan(input)
				pos += n

				switch token {
				case LAMBDA, LPAREN, RPAREN:
					return nil, pos - 1, UnexpectedToken
				case DOT:
					if !zn.IsRoot() {
						if _, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
							break definition
						}
					}

					return nil, pos - 1, UnexpectedToken
				case EOF:
					return nil, pos, UnexpectedEndOfInput
				default:
					zn.expr = &Abstraction{&Variable{string(token)}, nil}
					zn = zn.NewAbstractionZone()
				}
			}
		case DOT:
			return nil, pos - 1, UnexpectedToken
		case EOF:
			if zn.IsEmpty() {
				return nil, pos, UnexpectedEndOfInput
			}

			zn, _ = closeAbstractions(zn)

			if zn.IsEmpty() || !zn.IsRoot() {
				return nil, pos, UnexpectedEndOfInput
			}

			return zn.expr, pos, nil
		default:
			zn.SetOrApply(&Variable{string(token)})
		}
	}
}

func closeAbstractions(zn *zone) (*zone, Λ) {
	expr := zn.expr

	for !zn.IsRoot() {
		if abstr, isAbstr := zn.zn.expr.(*Abstraction); isAbstr {
			zn = zn.zn
			abstr.Body = expr
			expr = zn.expr
		} else {
			break
		}
	}

	return zn, expr
}

type zone struct {
	zn    *zone
	paren int
	expr  Λ
}

func newRootZone() *zone {
	return &zone{nil, 0, nil}
}

func (zn *zone) NewParenZone() *zone {
	return &zone{zn, zn.paren + 1, nil}
}

func (zn *zone) NewAbstractionZone() *zone {
	return &zone{zn, zn.paren, nil}
}

func (zn *zone) IsEmpty() bool {
	return zn.expr == nil
}

func (zn *zone) IsRoot() bool {
	return zn.zn == nil
}

func (zn *zone) SetOrApply(expr Λ) {
	if !zn.IsEmpty() {
		expr = &Application{zn.expr, expr}
	}

	zn.expr = expr
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
