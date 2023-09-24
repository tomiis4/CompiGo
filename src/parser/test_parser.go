package main

import "fmt"

const (
	NumberLiteral = "NumberLiteral"
	Symbol        = "Symbol"
)

type Token struct {
	Kind  string
	Value string
}

type Node struct {
	Operator Symbol
	Left     Node
	Right    Node
	Value    NumberLiteral
}

// TODO: add grammar like
// add: number + node|number
// multiply: number * node|number

func main() {
	// 5 + 4 * 2
	tokens := []Token{
		Token{Value: "5", Kind: NumberLiteral},
		Token{Value: "+", Kind: Symbol},
		Token{Value: "4", Kind: NumberLiteral},
		Token{Value: "*", Kind: Symbol},
		Token{Value: "2", Kind: NumberLiteral},
	}

	ast := Node{
		Operator: "+",
		Left: Node{
			Value: "5",
		},
		Right: Node{
			Operator: "*",
			Left: Node{
				Value: "4",
			},
			Right: Node{
				Value: "2",
			},
		},
	}
}
