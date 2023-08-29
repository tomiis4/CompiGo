package lexer

import "testing"

func BenchmarkLexer(b *testing.B) {
	content := []string{
		"print(\"Hello, world!\")",
		`main => {
		    print("Hello, world")
        }`,
	}

	for i := 0; i < len(content); i++ {
		for j := 0; j < b.N; j++ {
			_ = Lexer(content[i])
		}
	}
}
