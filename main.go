package main

import (
    "fmt"
	"io/ioutil"
    "os"
    "goop/scanner"
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
}

func readSource(sourceCodePath string) []byte {
    content, err := ioutil.ReadFile(sourceCodePath)
    if err != nil {
        fmt.Println(err)
    }
    return content
}

func printDebug(tokens []scanner.Token) {
    for _, token := range tokens {
        println(token.ToString())
    }
}
