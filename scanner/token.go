package scanner

type Token struct {
    tokenType   TokenType
    lexeme      string
    literal     string
    line        int
}

func (self *Token) ToString() string {
    return self.tokenType.String()
}
