package main

import (
    "fmt"
	"io/ioutil"
    "os"
    "goop/grammar"
    "goop/interpreter"
)

func main() {
    if len(os.Args) < 2 {
        os.Exit(64)
    }
    run(os.Args[1])
}

func run(sourcePath string) {
    source := readSource(sourcePath)

    println("parsing...")
    parser := grammar.NewParser(string(source))
    expr := parser.Parse()
    result := interpreter.Evaluate(expr)
    println(result.(string))
}

func readSource(sourceCodePath string) []byte {
    content, err := ioutil.ReadFile(sourceCodePath)
    if err != nil {
        fmt.Println(err)
    }
    return content
}

