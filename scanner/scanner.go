package scanner

const indentifierPattern string = "[a-zA-Z_][a-zA-Z_0-9]*"

var start = 0
var current = 0
var line = 1

type Scanner struct {
    SourceCode      string
    tokens          []Token
}

func (self *Scanner) ScanTokens() []Token {

    for !self.isAtEnd() {
        start = current;
        self.scanToken()
    }

    EOFToken := Token{EOF, "", "", line}
    self.tokens = append(self.tokens, EOFToken)

    return self.tokens
}



func (self *Scanner) scanToken() {
    c := self.SourceCode[current]
    current++

    switch c {
        case '(': self.addToken(LEFT_PAREN)
        case ')': self.addToken(RIGHT_PAREN)
        case '{': self.addToken(LEFT_BRACE)
        case '}': self.addToken(RIGHT_BRACE)
        case ',': self.addToken(COMMA)
        case '.': self.addToken(DOT)
        case '-': self.addToken(MINUS)
        case '+': self.addToken(PLUS)
        case ';': self.addToken(SEMICOLON)
        case '*': self.addToken(STAR)
        case '\n' : line++
    }
}

func (self *Scanner) addToken(params ...interface{}) {
    var literal string = ""
    tokenType := params[0].(TokenType)

    if  len(params) > 1 {
        literal = params[1].(string)
    }

    token := Token{
        tokenType: tokenType,
        lexeme: self.SourceCode[start:current],
        literal: literal,
        line: line,
    }
    self.tokens = append(self.tokens, token)
}

func (self *Scanner) isAtEnd() bool {
    return current >= len(self.SourceCode)
}

