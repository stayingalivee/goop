package scanner

import "goop/grammar"

var keywords = map[string]grammar.TokenType{
    "class":  grammar.CLASS,
    "else":   grammar.ELSE,
    "if":     grammar.IF,
    "flase":  grammar.FALSE,
    "true":   grammar.TRUE,
    "for":    grammar.FOR,
    "fun":    grammar.FUN,
    "nil":    grammar.NIL,
    "print":  grammar.PRINT,
    "return": grammar.RETURN,
    "super":  grammar.SUPER,
    "this":   grammar.THIS,
    "var":    grammar.VAR,
    "while":  grammar.WHILE,
}
