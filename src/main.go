package main

import (
    l "tomiis4/compigo/lexer"
)

func main() {
    content := map[string]string{
        "var": "var name: i32 = 610 // init variable",
        "function": "main => { /* comment */ }",
    }
    l.Lexer(content["var"])
}
