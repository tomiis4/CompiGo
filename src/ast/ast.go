package ast

import (
	"fmt"
	l "tomiis4/compigo/lexer"
)

type Declaration struct {
	Type       string
	Identifier string
	Assignment string
}

type Statement struct {
	Declarations Declaration
}

type Block struct {
	Statements []Statement
}

type Function struct {
	Name   string
	Return string
	Block  Block
}

type Tree struct {
	// imports
	// declarations
	Functions []Function
}

func Ast(tokens []l.Token) {
	fmt.Println("Hello, world")
	fmt.Println(tokens)
}

// ast.Program{
//         Functions: []ast.FunctionDeclaration{
//             {
//                 Name:       "main",
//                 ReturnType: "int",
//                 Block: ast.Block{
//                     Statements: []ast.Statement{
//                         ast.VariableDeclaration{
//                             Type:       "int",
//                             Identifier: "i",
//                         },
//                         ast.Assignment{
//                             Left:  ast.VariableReference{Name: "i"},
//                             Right: ast.IntegerLiteral{Value: 5},
//                         },
//                     },
//                 },
//             },
//         },
//     }
