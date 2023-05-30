package main

import (
    "fmt"
    "strconv"
    "strings"
)

// Token Types
const (
    INT = "INT"
    PLUS = "PLUS"
)

type Token struct {
    Type  string
    Value string
}

func Lex(input string) []Token {
    var tokens []Token

    for _, char := range strings.Split(input, " ") {
        switch {
        case char == "+":
            tokens = append(tokens, Token{PLUS, char})
        default:
            if _, err := strconv.Atoi(char); err == nil {
                tokens = append(tokens, Token{INT, char})
            }
        }
    }

    return tokens
}

func Parse(tokens []Token) int {
    result := 0

    for _, token := range tokens {
        switch token.Type {
        case INT:
            value, _ := strconv.Atoi(token.Value)
            result += value
        }
    }

    return result
}

func main() {
    input := "2 + 2 + 3 + 4"
    tokens := Lex(input)
    result := Parse(tokens)
    fmt.Println(result)
}