package lexer

import (
	"fmt"
	"regexp"
)

type Array = []string

var keywords = Array{
	"use", "var", "mut", "return", "type",
	"if", "elseif", "else", "for", "break",
	"contiune", "in"}
var symbols = Array{
	"(", ")", "[", "]", "{", "}", "<", ">",
	".", ",", "=", "==", "<=", ">=", "!=",
	"+=", "-=", "*=", "/=", "+", "-", "*",
	"/", "++", "--"}

const (
	Keyword       = "keyword"
	Identifier    = "identifier"
	Symbol        = "symbol"
	NumberLiteral = "NumberLiteral"
	stringLiteral = "StringLiteral"
)

type Token struct {
	kind  string
	value string
}

func removeComments(sPtr *string) *string {
    rLine := regexp.MustCompile(`\/\/.*$`)
    rMultiline := regexp.MustCompile(`(\/\*)(.*)(\*\/)`)

    s1 := rLine.ReplaceAllString(*sPtr, "")
    s2 := rMultiline.ReplaceAllString(s1, "")
    return &s2
}

/*

loop letter by letter

*/

// TODO: add error handling (wrong syntax)
func Lexer(content string) {
    content = *removeComments(&content)

	tokens := []Token{}
	currentToken := Token{}
    fmt.Println(content)
    fmt.Println(tokens, currentToken)

	// for i := 0; i < len(content); i++ {
	// 	char := string(content[i])
	// 	fmt.Println(char)
	// }
}
