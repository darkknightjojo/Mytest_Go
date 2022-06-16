package 底层原理

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func TestToken() {
	src := []byte("cos(x) + 2i * sin(x) // Euler")

	var s scanner.Scanner
	set := token.NewFileSet()
	file := set.AddFile("", set.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t\t%s\t%q\n", set.Position(pos), tok, lit)
	}
}
