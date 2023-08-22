package ast

import (
	"fmt"
	l "tomiis4/compigo/lexer"
)

type Block interface{}

type Statement struct {
	Type      string
	Condition string // FIXME
	Block     []Block
}

type Assignment struct {
	Left  []string
	Right []string
}

type Function struct {
	Name       string
	ReturnType string
	Block      []Block // assiegment, declaration, statement
}

type Import struct {
	Package string
	Name    string
}

type Declaration struct {
	Scope      string
	Type       string
	Identifier string
}

type Tree struct {
	Imports      []Import
	Declarations []Declaration
	Functions    []Function
}

func Ast(tokens []l.Token) {
	fmt.Println(tokens)

    tree := Tree{
        Functions: []Function{{
            Name: "main",
            ReturnType: "nil",
            Block: []Block{
                Declaration{
                    Scope: "main",
                    Type: "<string>",
                    Identifier: "arr",
                },
                Assignment{
                    Left: []string{"arr"},
                    Right: []string{"Hello", "world"},
                },
            },
        }},
    }
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
