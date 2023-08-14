package lexer

import (
	"regexp"
	"slices"
)

type Array = []string

var (
	keywords = Array{
		"use", "var", "mut", "return", "type",
		"if", "elseif", "else", "for", "break",
		"contiune", "in",
	}

	symbols = Array{
		"(", ")", "[", "]", "{", "}", "<", ">",
		".", ",", "=", "==", "<=", ">=", "!=",
		"+=", "-=", "*=", "/=", "=>", "+", "-",
        "*", "/", "++", "--", "#"
	}

	NumberRegex = regexp.MustCompile(`^-?\d+$`)
	StringRegex = regexp.MustCompile(`^".*"$`)
)

const (
	Keyword       = "keyword"
	Identifier    = "identifier"
	Symbol        = "symbol"
	NumberLiteral = "NumberLiteral"
	StringLiteral = "StringLiteral"
	Undefined     = "undefined"
)

type Token struct {
	Kind  string
	Value string
}

func getTokens(content string) []string {
	// REGEX: 1) "[^"]+"   == match string
	//        2) \w+       == match full words
	//        3) [^\s\w"]+ == match any character which is not space/word
	tokenRegex := regexp.MustCompile(`("[^"]*"|\w+|[^\s\w"]+)`)
	tokens := tokenRegex.FindAllString(content, -1)

	return tokens
}

func getKind(tk string) string {
	if slices.Contains(keywords, tk) {
		return Keyword
	} else if slices.Contains(symbols, tk) {
		return Symbol
	} else if NumberRegex.Match([]byte(tk)) {
		return NumberLiteral
	} else if StringRegex.Match([]byte(tk)) {
		return StringLiteral
	}

    return Identifier
}

// TODO: handle comments
func Lexer(content string) []Token {
	tokens := getTokens(content)
	tokenObjects := []Token{}

	for _, tk := range tokens {
		tokenObjects = append(tokenObjects, Token{
			Kind:  getKind(tk),
			Value: tk,
		})
	}
	return tokenObjects
}
