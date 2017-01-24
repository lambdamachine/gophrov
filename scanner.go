package λ

import (
	"bytes"
	"io"
)

type Scanner struct{}

func (scnr *Scanner) Scan(input io.RuneScanner) Token {
	for {
		r, _, err := input.ReadRune()

		if err != nil {
			return EOF
		}

		switch r {
		case ' ', '\n', '\t':
			continue
		default:
			switch r {
			case 'λ':
				return LAMBDA
			case '.':
				return DOT
			case '(':
				return LPAREN
			case ')':
				return RPAREN
			}

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
					buf.WriteRune(r)
				}
			}

		exit:
			input.UnreadRune()
			return Token(buf.String())
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
