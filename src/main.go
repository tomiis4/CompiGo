package main

import (
    "fmt"
    l "tomiis4/compigo/lexer"
)

func main() {
    content := map[string]string{
        "var": "var name: i32 = 610 // init variable",
        "function": "main => { \nprint(\"Test\")\n/*\n comment */\n }",
    }

    x := l.Lexer(content["function"])

    for _, v := range x {
        fmt.Println(v.Kind, v.Value) 
    }
}
