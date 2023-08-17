package main

import (
	// "fmt"
	l "tomiis4/compigo/lexer"
	a "tomiis4/compigo/ast"
)

func main() {
	content := map[string]string{
		"var": "var name: i32 = 610 // init variable",
		"function": `main => {
            arr := <string>["Hello", "world"]
            for k,v in arr {
                print(k,v)
            }
        }`,
	}

	tokens := l.Lexer(content["function"])

    a.Ast(tokens)
	// for _, v := range tokens {
	// 	fmt.Println(v.Kind, v.Value)
	// }
}
