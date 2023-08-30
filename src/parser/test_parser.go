package main

import "fmt"

const (
	NumberLiteral string
	Symbol        string
)

type Token struct {
	Kind  string
	Value string
}

type ExprInterface interface{}

type Expr struct {
    Left NumberLiteral
    Right ExprInterface // +-/*  ||  NumberLiteral 
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


    /*

    EXPRESSION: 

      +
     / \
    5  *
      / \
     4   2

    Expr   -> Expr + Term
        | Term

    Term   -> Term * Factor
        | Factor

    Factor -> NumberLiteral
        | ( Expr )

    */
}
