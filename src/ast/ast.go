package ast

import (
	"fmt"
	l "tomiis4/compigo/lexer"
)

// TODO
//  - add if statements
//  - add loops
//  - add assignments

type Block interface{}

type Value struct {
    Type string
    Value string
}

type Parameter struct {
    Name string
    Type string
}

type VarDeclaration struct {
    Name string
    Mutal bool
    Value []Value
}

type FuncCall struct {
    Name string
    Arguments []string
}

type FuncDeclaration struct {
    Name string
    Parameters []Parameter
    Body Block
}

type Import struct {
    Package string
    Name string
}

type Tree struct {
	Imports      []Import
    FuncDeclarations []FuncDeclaration
    
    //TODO: VarDeclarations []Declaration
}

// (function_declaration) ; [1:1 - 4:3]
//  name: (identifier) ; [1:10 - 13]
//  parameters: (parameters) ; [1:14 - 15]
//  body: (block) ; [2:5 - 3:24]
//   local_declaration: (variable_declaration) ; [2:5 - 24]
//    (assignment_statement) ; [2:11 - 24]
//     (variable_list) ; [2:11 - 11]
//      name: (identifier) ; [2:11 - 11]
//     (expression_list) ; [2:15 - 24]
//      value: (table_constructor) ; [2:15 - 24]
//       (field) ; [2:16 - 18]
//        value: (string) ; [2:16 - 18]
//         content: (string_content) ; [2:17 - 17]
//       (field) ; [2:21 - 23]
//        value: (string) ; [2:21 - 23]
//         content: (string_content) ; [2:22 - 22]
//   (function_call) ; [3:5 - 24]
//    name: (identifier) ; [3:5 - 9]
//    arguments: (arguments) ; [3:10 - 24]
//     (string) ; [3:11 - 23]
//      content: (string_content) ; [3:12 - 22]

func Ast(tokens []l.Token) {
    fmt.Println(tokens)
}
