package grammar

/*
Recursive descent is considered a top-down parser because it starts from the
top or outermost grammar rule (here  expression) and works its way down
into the nested subexpressions before finally reaching the leaves of the syntax
tree.

expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expresion ")" ;

    6 - 2 / 2
    6 / 2 - 2
    term
    factor - factor
    .
    .
    .
*/

type Parser struct {
    tokenList   []Token
    index       int
}

func NewParser(source string) *Parser {
    scnr := Scanner{SourceCode: string(source)}
    tokenList := scnr.ScanTokens()
    return &Parser{tokenList, 0}
}

func (self *Parser) Parse() {
    for !self.isAtEnd() {
        expr := self.Expression()
        tree := ""
        PrintTree(expr, &tree)
        println(tree)
    }
}

// ----------------
// parser controls
// ----------------
func (self *Parser) match(tokenTypeList ...TokenType) bool {
    for _, tokenType := range tokenTypeList {
        if !self.isAtEnd() && tokenType == self.peek().TokenType {

            self.next()
            return true
        }
    }
    return false
}

func (self *Parser) previous() *Token {
    return &self.tokenList[self.index - 1]
}

func (self *Parser) next() *Token {
    // look ahead by 1
    if !self.isAtEnd() {
        self.index++
    }
    return self.previous()
}

func (self *Parser) isAtEnd() bool {
    return self.peek().TokenType == EOF
}

func (self *Parser) peek() Token {
    return self.tokenList[self.index]
}


// --------------------------
// parser grammar rules impl
// --------------------------

// expression     → equality ;
func (self *Parser) Expression() *Expr {
    return self.Equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (self *Parser) Equality() *Expr {
    expr := self.Comparision()

    for self.match(BANG_EQUAL, EQUAL_EQUAL) {
        operator := &Operator{self.previous()}
        right := self.Comparision()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (self *Parser) Comparision() *Expr {
    expr := self.Term()

    for self.match(LESS, LESS_EQUAL, GREATER, GREATER_EQUAL) {
        operator := &Operator{self.previous()}
        right := self.Term()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

// term           → factor ( ( "-" | "+" ) factor )* ;
func (self *Parser) Term() *Expr {
    expr := self.Factor()

    for self.match(PLUS, MINUS) {
        operator := &Operator{self.previous()}
        right := self.Factor()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

// factor         → unary ( ( "/" | "*" ) unary )* ;
func (self *Parser) Factor() *Expr {
    expr := self.Unary()

    for self.match(SLASH, STAR) {
        operator := &Operator{self.previous()}
        right := self.Unary()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

// unary          → ( "!" | "-" ) unary | primary ;
func (self *Parser) Unary() *Expr {

    if self.match(BANG, MINUS) {
        operator := self.previous()
        right := self.Unary()
        expr := BuildUnaryExpr(operator, right)
        return expr
    }

    return self.Primary()
}

// primary        → NUMBER | STRING | "true" | "false" | "nil"  | "(" expresion ")" ;
func (self *Parser) Primary() *Expr {

    if self.match(NUMBER, STRING, TRUE, FALSE, NIL) {
        return BuildLiteralExpr(self.previous())
    }

    if self.match(LEFT_PAREN) {
        expr := self.Expression()

        // now a closing paranthesis is expected. otherwise it's syntax error
        if self.match(RIGHT_PAREN) {
            return BuildGroupingExpr(expr, self.peek().Line)
        }
    }

    panic("syntax error")
}

// expression data types
type Expr struct{
    Literal     *Literal
    Unary       *Unary
    Binary      *Binary
    Grouping    *Grouping
}

type Literal struct{
    Token       *Token
}

type Unary struct{
    Token       *Token
    Expr        *Expr
}

type Binary struct{
    Left        *Expr
    Operator    *Operator
    Right       *Expr
}

type Grouping struct{
    RightParan  *Token
    Expr        *Expr
    LeftParan   *Token
}

type Operator struct{
    Token       *Token
}

// using builder functions cuz I genuinely hate typing up struct inits
// inside other struct inits, it's ugly af like seriously super ugly
func BuildLiteralExpr(token *Token) *Expr {
    return &Expr{Literal: &Literal{token}}
}

func BuildUnaryExpr(token *Token, expr *Expr) *Expr {
    return &Expr{Unary: &Unary{token, expr}}
}

func BuildBinaryExpr(left *Expr, operator *Operator, right *Expr) *Expr {
    return &Expr{Binary: &Binary{left, operator, right}}
}

func BuildGroupingExpr(expr *Expr, line int) *Expr {
    rightParan := NewToken(RIGHT_PAREN, "(", "", line)
    leftParan  := NewToken(LEFT_PAREN, ")", "", line)
    return &Expr{Grouping: &Grouping{rightParan, expr, leftParan}}
}

func NewOperator(token *Token) *Operator {
    return &Operator{Token: token}
}

