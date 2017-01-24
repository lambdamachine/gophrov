package λ

import (
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, error) {
	var scnr Scanner

	for {
		token := scnr.Scan(input)

		switch token {
		case EOF, LAMBDA, DOT, LPAREN, RPAREN:
		default:
			return &Variable{string(token)}, nil
		}
	}
}
