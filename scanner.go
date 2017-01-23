package λ

import "io"

type Scanner struct {
}

func NewScanner(input io.Reader) *Scanner {
	return &Scanner{}
}

func (scnr *Scanner) Scan() Token {
	return nil
}
