package grammar



func PrintTree(node interface{}, tree *string) {

    switch node.(type) {

    case *Literal:
            *tree += node.(*Literal).Token.Lexeme
            return

        case *Operator:
            Operator := node.(*Operator)
            *tree += Operator.Token.Lexeme
            return

        case *Unary:
            unary := node.(*Unary)
            *tree += unary.Token.Lexeme
            PrintTree(unary.Expr, tree)

        case *Binary:
            binary := node.(*Binary)
            PrintTree(binary.Operator, tree)
            PrintTree(binary.Left, tree)
            PrintTree(binary.Right, tree)

        case *Grouping:
            grouping := node.(*Grouping)
            *tree += grouping.RightParan.Lexeme
            PrintTree(grouping.Expr, tree)
            *tree += grouping.LeftParan.Lexeme

        case *Expr:
            Expr := node.(*Expr)
            if Expr.Binary != nil {
                PrintTree(Expr.Binary, tree)
            } else if Expr.Unary != nil {
                PrintTree(Expr.Unary, tree)
            } else if Expr.Grouping != nil {
                PrintTree(Expr.Grouping, tree)
            } else {
                PrintTree(Expr.Literal, tree)
            }
    }
}
