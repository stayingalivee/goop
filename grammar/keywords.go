package grammar


var keywords = map[string]TokenType{
    "class":  CLASS,
    "else":   ELSE,
    "if":     IF,
    "flase":  FALSE,
    "true":   TRUE,
    "for":    FOR,
    "fun":    FUN,
    "nil":    NIL,
    "print":  PRINT,
    "return": RETURN,
    "super":  SUPER,
    "this":   THIS,
    "var":    VAR,
    "while":  WHILE,
}
