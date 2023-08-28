package ast

import (
	"fmt"
	l "tomiis4/compigo/lexer"
)

// TODO
//  - add loops
//  - add assignment (objects)

type Block interface{}

type Condition interface{}

type Alternative interface{}

type Value struct {
	Type  string
	Value string
}

type ConditionValue struct {
	Left   string
	Middle string
	Right  string
}

type ConditionSide struct {
	Left  []Condition
	Right []Condition
}

type ElseIfStatement struct {
	Condition Condition
	Body      Block
}

type ElseStatement struct {
	Body Block
}

type IfStatement struct {
	Condition   Condition
	Body        Block
	Alternative Alternative
}

type Parameter struct {
	Name string
	Type string
}

type VarDeclaration struct {
	Name  string
	Mutal bool
	Value []Value
}

type FuncCall struct {
	Name      string
	Arguments []string
}

type FuncDeclaration struct {
	Name       string
	Parameters []Parameter
	Body       Block
}

type Import struct {
	Package string
	Name    string
}

type Tree struct {
	Imports          []Import
	FuncDeclarations []FuncDeclaration
	VarDeclarations  []VarDeclaration
}

func Ast(tokens []l.Token) {
	fmt.Println(tokens)
}
