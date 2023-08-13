package lexer

import (
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
	Kind  string
	Value string
}

func removeComments(input string) string {
	lineRegex := regexp.MustCompile(`\/\/.*$`)
	multilineRegex := regexp.MustCompile(`\/\*.*?\*\/`)

	return multilineRegex.ReplaceAllString(
		lineRegex.ReplaceAllString(input, ""),
		"")
}

func getTokens(content string) []string {
    // FIXME: does not split characters, example: function{
	tokenRegex := regexp.MustCompile(`("[^"]*"|[^\s"]+)`)
	tokens := tokenRegex.FindAllString(content, -1)

	return tokens
}

func Lexer(content string) []Token {
	content = removeComments(content)

	tokens := getTokens(content)
	tokenObjects := []Token{}

	for _, tk := range tokens {
		tokenObjects = append(tokenObjects, Token{
			Kind:  "TODO",
			Value: tk,
		})
	}
	return tokenObjects
}
