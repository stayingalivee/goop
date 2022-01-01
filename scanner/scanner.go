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
    c := self.next()
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
        case '!': self.addTokenOnCondition(self.matchNext('='), BANG_EQUAL, BANG)
        case '=': self.addTokenOnCondition(self.matchNext('='), EQUAL_EQUAL, EQUAL)
        case '>': self.addTokenOnCondition(self.matchNext('='), GREATER_EQUAL, GREATER)
        case '<': self.addTokenOnCondition(self.matchNext('='), LESS_EQUAL, LESS)
        case '/':
            if self.matchNext('/') {
                for self.peek() != '\n' && !self.isAtEnd() {
                    self.next()
                }
            } else {
                self.addToken(SLASH)
            }
        case ' ':
        case '\t':
        case '\r':
        case '\n': line++
        default: println("illegal char")
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

func (self *Scanner) addTokenOnCondition(
    condition bool, ifTrue TokenType, ifFalse TokenType) {
    if condition {
        self.addToken(ifTrue)
    } else {
        self.addToken(ifFalse)
    }
}

func (self *Scanner) isAtEnd() bool {
    return current >= len(self.SourceCode)
}

func (self *Scanner) next() byte {
    current++;
    return self.SourceCode[current - 1]
}

func (self * Scanner) matchNext(expected byte) bool {
    if self.isAtEnd() {
        return false
    }
    if(self.SourceCode[current] != expected) {
        return false
    }
    current++
    return true
}

func (self *Scanner) peek() byte{
    return self.SourceCode[current]
}
