package λ

import (
	"fmt"
	"io"
)

type Parser struct{}

func (prsr *Parser) Parse(input io.RuneScanner) (Λ, error) {
	return nil, fmt.Errorf("not implemented yet")
}
