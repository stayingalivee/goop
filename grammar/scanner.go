package grammar

import "os"

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
    switch  {
        case c == '(': self.addToken(LEFT_PAREN)
        case c == ')': self.addToken(RIGHT_PAREN)
        case c == '{': self.addToken(LEFT_BRACE)
        case c == '}': self.addToken(RIGHT_BRACE)
        case c == ',': self.addToken(COMMA)
        case c == '-': self.addToken(MINUS)
        case c == '+': self.addToken(PLUS)
        case c == ';': self.addToken(SEMICOLON)
        case c == '*': self.addToken(STAR)
        case c == '!': self.addTokenOnCondition(self.matchNext('='), BANG_EQUAL, BANG)
        case c == '=': self.addTokenOnCondition(self.matchNext('='), EQUAL_EQUAL, EQUAL)
        case c == '>': self.addTokenOnCondition(self.matchNext('='), GREATER_EQUAL, GREATER)
        case c == '<': self.addTokenOnCondition(self.matchNext('='), LESS_EQUAL, LESS)
        case c == ' ':  // ignore all whitespaces 
        case c == '\t':
        case c == '\r':
        case c == '\n': line++
        case c == '.':
            if self.isNumeric(self.peek()) {
                self.handleNumber()
            } else {
                self.addToken(DOT)
            }
        case c == '/':
            if self.matchNext('/') {
                for self.peek() != '\n' && !self.isAtEnd() {
                    self.next()
                }
            } else {
                self.addToken(SLASH)
            }
        case c == '|':
            if self.matchNext('|') {
                self.addToken(OR)
            }
        case c== '&':
            if self.matchNext('&') {
                self.addToken(AND)
            }
        case c == '"': self.handleString()
        case self.isNumeric(c): self.handleNumber()

        default:
            if self.isNumeric(c) {
                self.handleNumber()
            } else if self.isAlpha(c) {
                self.handleIdentifier()
            } else {
                print("Illegal syntax")
            }
    }
}

func (self *Scanner) addToken(params ...interface{}) {
    var literal string = ""
    tokenType := params[0].(TokenType)

    if  len(params) > 1 {
        literal = params[1].(string)
    }

    token := Token{
        TokenType: tokenType,
        Lexeme: self.SourceCode[start:current],
        Literal: literal,
        Line: line,
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

func (self *Scanner) handleString() {
    for self.peek() != '"' && !self.isAtEnd() {
        if self.peek() == '\n' {
            line++
        }
        self.next()
    }

    if self.isAtEnd() {
        print("Error: Unterminated String")
    }
    self.next()
    literal := self.SourceCode[start + 1: current - 1]
    self.addToken(STRING, literal)
}

var dotCount = 0;
func (self *Scanner) isNumeric(c byte) bool {
    if c == '.' {
        dotCount++
    }
    if dotCount > 1 {
        print("Error: illegal number format")
        os.Exit(64)
    }
    return (c >= '0' && c <= '9') || c == '.'
}

func (self *Scanner) handleNumber() {
    for self.isNumeric(self.peek()) && !self.isAtEnd() {
        self.next()
    }
    dotCount = 0
    literal := self.SourceCode[start: current]
    self.addToken(NUMBER, literal)
}

func (self *Scanner) handleIdentifier() {
    for self.isAlphaNumeric(self.peek()) {
        self.next()
    }
    lexeme := self.SourceCode[start: current]

    // check if reserved keyword
    tokenType, exists := keywords[lexeme]
    if exists {
        self.addToken(tokenType)
    } else {
        self.addToken(IDENTIFIER)
    }
}

func (self *Scanner) isAlpha(c byte) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
            c == '_'
}

func (self *Scanner) isAlphaNumeric(c byte) bool {
    return self.isAlpha(c) || self.isNumeric(c)
}

