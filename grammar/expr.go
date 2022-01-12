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

