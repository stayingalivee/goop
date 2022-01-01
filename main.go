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
    readSource(os.Args[1])
}

func run(source string) {
    // TODO: implement compiler lmao
}

func readSource(sourceCodePath string) {
    content, err := ioutil.ReadFile(sourceCodePath)
    if err != nil {
        fmt.Println(err)
    }

    scnr := scanner.Scanner{
        SourceCode: string(content),
    }

    tokens := scnr.ScanTokens()
    for _, token := range tokens {
        print(token.ToString() + ", ")
    }
}

