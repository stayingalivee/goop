package grammar

import "goop/scanner"

/*
expression     → literal
               | unary
               | binary
               | grouping ;
literal        → NUMBER | STRING | "true" | "false" | "nil" ;
grouping       → "(" expression ")" ;
unary          → ( "-" | "!" ) expression ;
binary         → expression operator expression ;
operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
               | "+"  | "-"  | "*" | "/" ;
*/

type Expr struct{
    Literal     *Literal
    Unary       *Unary
    Binary      *Binary
    Grouping    *Grouping
}

type Literal struct{
    Token       *scanner.Token
}

type Unary struct{
    Token       *scanner.Token
    Expr        *Expr
}

type Binary struct{
    Left        *Expr
    Operator    *Operator
    Right       *Expr
}

type Grouping struct{
    RightParan  *scanner.Token
    Expr        *Expr
    LeftParan   *scanner.Token
}

type Operator struct{
    Token       *scanner.Token
}

func BuildLiteralExpr(token *scanner.Token) *Expr {
    return &Expr{Literal: &Literal{token}}
}

func BuildUnaryExpr(token *scanner.Token, expr *Expr) *Expr {
    return &Expr{Unary: &Unary{token, expr}}
}

func BuildBinaryExpr(left *Expr, operator *Operator, right *Expr) *Expr {
    return &Expr{Binary: &Binary{left, operator, right}}
}

func BuildGroupingExpr(expr *Expr, line int) *Expr {
    rightParan := &scanner.Token{scanner.RIGHT_PAREN, "(", "", line}
    leftParan  := &scanner.Token{scanner.LEFT_PAREN, ")", "", line}
    return &Expr{Grouping: &Grouping{rightParan, expr, leftParan}}
}

