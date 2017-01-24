package λ

import (
	"bytes"
	"io"
)

type Scanner struct{}

func (scnr *Scanner) Scan(input io.RuneScanner) (tok Token, n int) {
	for {
		r, _, err := input.ReadRune()

		if err != nil {
			tok = EOF
			return
		}

		n++

		switch r {
		case ' ', '\n', '\t':
			continue
		default:
			switch r {
			case 'λ':
				tok = LAMBDA
			case '.':
				tok = DOT
			case '(':
				tok = LPAREN
			case ')':
				tok = RPAREN
			default:
				var tn int
				input.UnreadRune()
				tok, tn = scanVariable(input)
				n += tn - 1 // reduce an unread rune from 2-lines above...
			}

			return
		}
	}
}

func scanVariable(input io.RuneScanner) (tok Token, n int) {
	var buf bytes.Buffer

	for {
		r, _, err := input.ReadRune()

		if err != nil {
			goto exit
		}

		switch r {
		case ' ', '\n', '\t', 'λ', '.', '(', ')':
			goto exit
		default:
			n++
			buf.WriteRune(r)
		}
	}

exit:
	input.UnreadRune()
	tok = Token(buf.String())
	return
}

type Token string

const (
	EOF    Token = ""
	LAMBDA Token = "λ"
	DOT    Token = "."
	LPAREN Token = "("
	RPAREN Token = ")"
)
