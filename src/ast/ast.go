package ast

import (
    "fmt"
	l "tomiis4/compigo/lexer"
)

func Ast(tokens []l.Token) {
    fmt.Println("Hello, world")
    fmt.Println(tokens)
}
