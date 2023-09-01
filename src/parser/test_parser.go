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

    temp_node := []Node{}
    for _, v := range tokens {
        if temp_node.Left == nil && v.Kind == NumberLiteral {
            temp_node.Left = v.Value
        } else if temp_node.Right == nil && v.Kind == NumberLiteral {
            temp_node.Right = v.Value
        }
        // } else if temp_node.Operator == 
    }
}
