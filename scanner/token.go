package scanner

type Token struct {
    TokenType   TokenType
    Lexeme      string
    Literal     string
    Line        int
}

func (self *Token) ToString() string {
    return "type: " + self.TokenType.String() +
           ", lexeme: " + self.Lexeme +
           ", literal " + self.Literal
}
