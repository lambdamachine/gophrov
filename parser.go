package λ

import (
	"errors"
	"io"
)

type Parser struct {
	Report Reporter
	scnr   Scanner
}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, int, error) {
	var (
		pos = 0
		zn  = newRootZone()
	)

	if nil == prsr.Report {
		prsr.Report = func(_ Report) error {
			return nil
		}
	}

	prsr.Report(nil)

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
				zn = closeAbstractions(zn)

				if zn.IsEmpty() || zn.IsRoot() {
					return nil, pos - 1, UnexpectedToken
				}

				expr := zn.expr
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

			zn = closeAbstractions(zn)

			if zn.IsEmpty() || !zn.IsRoot() {
				return nil, pos, UnexpectedEndOfInput
			}

			return zn.expr, pos, nil
		default:
			zn.SetOrApply(&Variable{string(token)})
		}
	}
}

func closeAbstractions(zn *zone) *zone {
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

	return zn
}

type Reporter func(Report) error

type Report interface {
	Event() ParserEvent
	Expr() Λ
	Pos() int
}

type ParserEvent int

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

var (
	UnexpectedEndOfInput = errors.New("unexpected end of input")
	UnexpectedToken      = errors.New("unexpected token")
)
