package λ

import (
	"bytes"
	"io"
)

type Scanner struct{}

func (scnr *Scanner) Scan(input io.RuneScanner) (tok Token, pos int) {
	for {
		r, _, err := input.ReadRune()

		if err != nil {
			tok = EOF
			return
		}

		pos++

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
				goto variable
			}

			return

		variable:
			var buf bytes.Buffer
			buf.WriteRune(r)

			for {
				r, _, err := input.ReadRune()

				if err != nil {
					goto exit
				}

				switch r {
				case ' ', '\n', '\t', 'λ', '.', '(', ')':
					goto exit
				default:
					pos++
					buf.WriteRune(r)
				}
			}

		exit:
			input.UnreadRune()
			tok = Token(buf.String())
			return
		}
	}
}

type Token string

const (
	EOF    Token = ""
	LAMBDA Token = "λ"
	DOT    Token = "."
	LPAREN Token = "("
	RPAREN Token = ")"
)
