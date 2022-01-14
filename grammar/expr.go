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

func New(tokenList []Token) *Parser {
    return &Parser{tokenList, 0}
}

func (parser *Parser) Parse() *Expr {
    //TODO : call stuff
}

func (self *Parser) match(tokenTypeList ...TokenType) bool {
    for _, tokenType := range tokenTypeList {
        if tokenType == self.tokenList[self.index].TokenType {
            return true
        }
    }
    return false
}

func (self *Parser) previous() {
    //TODO: previous is used as if its returning an operator\
    //but maybe we need to return a token instead.
}

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


// expression     → equality ;
func (self *Parser) Expression() *Expr {
    return self.Equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (self *Parser) Equality() *Expr {
    expr := self.Comparision()

    for self.match(BANG_EQUAL, EQUAL_EQUAL) {
        opeartor := self.previous()
        right := self.Comparision()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (self *Parser) Comparision() *Expr {
    expr := self.Term()

    for self.match(LESS, LESS_EQUAL, GREATER, GREATER_EQUAL) {
        operator := self.previous()
        right := self.Term()
        expr = BuildBinaryExpr(expr, operator, right)
    }
    return expr
}

