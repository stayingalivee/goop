package main

import (
    "fmt"
	"io/ioutil"
    "os"
    "goop/scanner"
    "goop/grammar"
)

func main() {
    if len(os.Args) < 2 {
        os.Exit(64)
    }
    run(os.Args[1])
}

func run(sourcePath string) {
    content := readSource(sourcePath)

    scnr := scanner.Scanner{SourceCode: string(content)}
    tokens := scnr.ScanTokens()

    printDebug(tokens)

    head := getTree()
    printAstTree(head)
}

func readSource(sourceCodePath string) []byte {
    content, err := ioutil.ReadFile(sourceCodePath)
    if err != nil {
        fmt.Println(err)
    }
    return content
}

func printDebug(tokens []scanner.Token) {
    println("tokenization..")
    for _, token := range tokens {
        println(token.ToString())
    }
}

func printAstTree(node interface{}) {
    println("printing ast")
    tree := ""
    grammar.PrintTree(node, &tree)
    println(tree)
}

func getTree() *grammar.Expr {

    literalExpr := grammar.BuildLiteralExpr(
        scanner.NewToken(scanner.NUMBER, "123", "", 1))
    unaryExpr := grammar.BuildUnaryExpr(
        scanner.NewToken(scanner.MINUS, "-", "", 1),
        literalExpr,
    )

    star := grammar.NewOperator(scanner.NewToken(scanner.STAR, "*", "", 1))

    groupedExpr := grammar.BuildLiteralExpr(scanner.NewToken(scanner.NUMBER, "13.21", "", 1))
    grouping := grammar.BuildGroupingExpr(groupedExpr, 1)

    binary := grammar.BuildBinaryExpr(unaryExpr, star, grouping)

    return binary
}

