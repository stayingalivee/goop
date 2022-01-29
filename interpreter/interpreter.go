package interpreter

import (
	"fmt"
	"goop/grammar"
	"strconv"
)

// -1


func Evaluate(expr *grammar.Expr) interface{} {

    if expr.Literal != nil {
            return handleLiteral(expr.Literal)
    } else if expr.Unary != nil {
            return handleUnary(expr.Unary)
    } else if expr.Binary != nil {
            return handleBinary(expr.Binary)
    } else if expr.Grouping != nil {
            return handleGrouping(expr.Grouping)
    }
    // unreachable code
    panic("you should not see this error")
}

func handleLiteral(literal *grammar.Literal) string {
    return literal.Token.Literal
}

func handleUnary(unary *grammar.Unary) interface{} {
    switch unary.Token.TokenType {
        case grammar.MINUS:
            result, err := strconv.Atoi(Evaluate(unary.Expr).(string))
            if err!= nil {
                return -result
            }
        case grammar.BANG:
            result, err := strconv.ParseBool(Evaluate(unary.Expr).(string))
            if err!= nil {
                return !result
            }
    }
    panic("panic lan")
}

func handleBinary(binary *grammar.Binary) interface{} {


    switch binary.Operator.Token.TokenType {
        case grammar.MINUS:
            fallthrough
        case grammar.PLUS:
            fallthrough
        case grammar.STAR:
            fallthrough
        case grammar.SLASH:
            return handleArithmaticBinary(binary)
        case grammar.AND:
            fallthrough
        case grammar.OR:
            return handleLogicalBinary(binary)
        case grammar.EQUAL:
            fallthrough
        case grammar.BANG_EQUAL:
            fallthrough
        case grammar.LESS:
            fallthrough
        case grammar.LESS_EQUAL:
            fallthrough
        case grammar.GREATER:
            fallthrough
        case grammar.GREATER_EQUAL:
            return handleComparisionBinary(binary)
    }
    panic("panic lan")
}

func handleArithmaticBinary(binary *grammar.Binary) string{
    left, _ := strconv.Atoi(Evaluate(binary.Left).(string))
    right, _ :=strconv.Atoi(Evaluate(binary.Right).(string))
    println(left)
    println(right)

    switch binary.Operator.Token.TokenType {
        case grammar.MINUS:
            return fmt.Sprint(left - right)
        case grammar.PLUS:
            return fmt.Sprint(left + right)
        case grammar.STAR:
            return fmt.Sprint(left * right)
        case grammar.SLASH:
            return fmt.Sprint(left / right)
    }
    panic("arithmatic error")
}

func handleLogicalBinary(binary *grammar.Binary) bool {
    left, _ := strconv.ParseBool(Evaluate(binary.Left).(string))
    right, _ :=strconv.ParseBool(Evaluate(binary.Right).(string))

    switch binary.Operator.Token.TokenType {
        case grammar.AND:
            return left && right
        case grammar.OR:
            return left || right
    }
    panic("bool error")
}

func handleComparisionBinary(binary *grammar.Binary) bool {
    left, _ := strconv.Atoi(Evaluate(binary.Left).(string))
    right, _ :=strconv.Atoi(Evaluate(binary.Right).(string))

    switch binary.Operator.Token.TokenType {
        case grammar.EQUAL:
            return left == right
        case grammar.BANG_EQUAL:
            return left != right
        case grammar.LESS:
            return left < right
        case grammar.LESS_EQUAL:
            return left <= right
        case grammar.GREATER:
            return left > right
        case grammar.GREATER_EQUAL:
            return left >= right
    }
    panic("comparision error")
}

// 6 * (1 + 2)
func handleGrouping(grouping *grammar.Grouping) interface{} {
    return Evaluate(grouping.Expr)
}
