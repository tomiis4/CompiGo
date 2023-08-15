package lexer

import (
	// "reflect"
	"testing"
)

func tok(kind, value string) Token {
	return Token{Kind: kind, Value: value}
}

func compare(lexerOutput, expectedOutput []Token, t *testing.T) {
	lenLexer := len(lexerOutput)
	lenExpected := len(expectedOutput)

	if lenLexer != lenExpected {
		t.Errorf("incorrect size: output = %d, result = %d", lenLexer, lenExpected)
	}

	for i := 0; i < lenLexer; i++ {
		outputToken := lexerOutput[i]
		expectedToken := expectedOutput[i]

		if outputToken.Kind != expectedToken.Kind {
			t.Errorf("KIND: expected %v and got %v, index %d",
				expectedToken, outputToken, i)
		}

		if outputToken.Value != expectedToken.Value {
			t.Errorf("Value: expected %v and got %v, index %d",
				expectedToken.Value, outputToken.Value, i)
		}
	}
}

func TestLexerFunction(t *testing.T) {
	s := `main => {
        print("Hello, world")
    }`

	lexerOutput := Lexer(s)
	expectedOutput := []Token{
		tok(Identifier, "main"),
		tok(Symbol, "=>"),
		tok(Symbol, "{"),
		tok(Identifier, "print"),
		tok(Symbol, "("),
		tok(StringLiteral, "\"Hello, world\""),
		tok(Symbol, ")"),
		tok(Symbol, "}"),
	}

	compare(lexerOutput, expectedOutput, t)
}

func TestLexerVariables(t *testing.T) {
	s := `var x: i8 = 45
    var y: f32
    arr := <string, 5>["Ahoj"]
    `

	lexerOutput := Lexer(s)
	expectedOutput := []Token{
		tok(Keyword, "var"),
		tok(Identifier, "x"),
		tok(Symbol, ":"),
		tok(Identifier, "i8"),
		tok(Symbol, "="),
		tok(NumberLiteral, "45"),
		tok(Keyword, "var"),
		tok(Identifier, "y"), tok(Symbol, ":"),
		tok(Identifier, "f32"),
		tok(Identifier, "arr"),
		tok(Symbol, ":="),
		tok(Symbol, "<"),
		tok(Identifier, "string"),
		tok(Symbol, ","),
		tok(NumberLiteral, "5"),
		tok(Symbol, ">"),
		tok(Symbol, "["),
		tok(StringLiteral, "\"Ahoj\""),
		tok(Symbol, "]"),
	}

	compare(lexerOutput, expectedOutput, t)
}

func TestLexerCombined(t *testing.T) {
	s := `main => {
        arr := <string, 5>["Ahoj"]
        print("Hello, world", arr[0])
    }`

	lexerOutput := Lexer(s)
	expectedOutput := []Token{
		tok(Identifier, "main"),
		tok(Symbol, "=>"),
		tok(Symbol, "{"),
		tok(Identifier, "arr"),
		tok(Symbol, ":="),
		tok(Symbol, "<"),
		tok(Identifier, "string"),
		tok(Symbol, ","),
		tok(NumberLiteral, "5"),
		tok(Symbol, ">"),
		tok(Symbol, "["),
		tok(StringLiteral, "\"Ahoj\""),
		tok(Symbol, "]"),
		tok(Identifier, "print"),
		tok(Symbol, "("),
		tok(StringLiteral, "\"Hello, world\""),
		tok(Symbol, ","),
		tok(Identifier, "arr"),
		tok(Symbol, "["),
		tok(NumberLiteral, "0"),
		tok(Symbol, "]"),
		tok(Symbol, ")"),
		tok(Symbol, "}"),
	}

	compare(lexerOutput, expectedOutput, t)
}
