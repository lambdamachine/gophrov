package λ

import "io"

type Scanner struct {
}

func NewScanner(input io.Reader) *Scanner {
	return &Scanner{}
}

func (scnr *Scanner) Scan() Token {
	return EOF
}

type Token string

const (
	EOF    Token = ""
	LAMBDA Token = "λ"
	DOT    Token = "."
)
