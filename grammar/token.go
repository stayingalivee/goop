package grammar

type Token struct {
    TokenType   TokenType
    Lexeme      string
    Literal     string
    Line        int
}

func NewToken(tokenType TokenType,
    lexeme string, literal string, line int) *Token {
    return &Token{
        TokenType: tokenType,
        Lexeme:    lexeme,
        Literal:   literal,
        Line:      line,
    }
}

func (self *Token) ToString() string {
    return "type: " + self.TokenType.String() +
           ", lexeme: " + self.Lexeme +
           ", literal " + self.Literal
}
