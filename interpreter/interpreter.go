package interpreter

import (
	"fmt"
	"goop/grammar"
	"strconv"
)

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

    evalUnary := Evaluate(unary.Expr).(string)
    switch unary.Token.TokenType {
        case grammar.MINUS:
            result, _ := strconv.Atoi(evalUnary)
            return -result
        case grammar.BANG:
            result, _ := strconv.ParseBool(evalUnary)
            return !result
    }
    // unreachable code
    panic("you should not see this error")
}

func handleBinary(binary *grammar.Binary) interface{} {

    switch binary.Operator.Token.TokenType {

        case grammar.MINUS,
             grammar.PLUS,
             grammar.STAR,
             grammar.SLASH:
            return handleArithmeticBinary(binary)

        case grammar.AND,
             grammar.OR:
            return handleLogicalBinary(binary)

        case grammar.EQUAL,
             grammar.BANG_EQUAL,
             grammar.LESS,
             grammar.LESS_EQUAL,
             grammar.GREATER,
             grammar.GREATER_EQUAL:
            return handleComparisionBinary(binary)
    }

    // unreachable code
    panic("you should not see this error")
}

func handleArithmeticBinary(binary *grammar.Binary) string {
    left, _ := strconv.Atoi(Evaluate(binary.Left).(string))
    right, _ :=strconv.Atoi(Evaluate(binary.Right).(string))

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
    panic("arithmetic error")
}

func handleLogicalBinary(binary *grammar.Binary) string {
    left, _ := strconv.ParseBool(Evaluate(binary.Left).(string))
    right, _ :=strconv.ParseBool(Evaluate(binary.Right).(string))
    println("left")
    println(left)
    println("right")
    println(right)

    switch binary.Operator.Token.TokenType {
        case grammar.AND:
            return fmt.Sprint(left && right)
        case grammar.OR:
            return fmt.Sprint(left || right)
    }
    panic("bool error")
}

func handleComparisionBinary(binary *grammar.Binary) string {

    left, _ := strconv.Atoi(Evaluate(binary.Left).(string))
    right, _ :=strconv.Atoi(Evaluate(binary.Right).(string))

    switch binary.Operator.Token.TokenType {
        case grammar.EQUAL:
            return fmt.Sprint(left == right)
        case grammar.BANG_EQUAL:
            return fmt.Sprint(left != right)
        case grammar.LESS:
            return fmt.Sprint(left < right)
        case grammar.LESS_EQUAL:
            return fmt.Sprint(left <= right)
        case grammar.GREATER:
            return fmt.Sprint(left > right)
        case grammar.GREATER_EQUAL:
            return fmt.Sprint(left >= right)
    }
    panic("comparision error")
}

// 6 * (1 + 2)
func handleGrouping(grouping *grammar.Grouping) interface{} {
    return Evaluate(grouping.Expr)
}
