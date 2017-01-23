package λ

import (
	"bufio"
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
			}

			return EOF
		}
	}
}

type Token string

const (
	EOF    Token = ""
	LAMBDA Token = "λ"
	DOT    Token = "."
)
