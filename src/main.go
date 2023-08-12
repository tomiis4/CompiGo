package main

import (
    "fmt"
    l "tomiis4/compigo/lexer"
)

func main() {
    content := map[string]string{
        "var": "var name: i32 = 610 // init variable",
        "function": "main => { /* comment */ }",
    }
    fmt.Println(l.Lexer(content["function"]))
}
