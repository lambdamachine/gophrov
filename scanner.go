package λ

import (
	"bufio"
	"bytes"
	"io"
)

type Scanner struct {
	input io.RuneScanner
}

func NewScanner(input io.Reader) *Scanner {
	return &Scanner{input: bufio.NewReader(input)}
}

func (scnr *Scanner) Scan() Token {
	for {
		r, _, err := scnr.input.ReadRune()

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
				r, _, err := scnr.input.ReadRune()

				if err != nil {
					scnr.input.UnreadRune()
					return Token(buf.String())
				}

				switch r {
				case ' ', '\n', '\t', 'λ', '.', '(', ')':
					scnr.input.UnreadRune()
					return Token(buf.String())
				default:
					buf.WriteRune(r)
				}
			}
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
