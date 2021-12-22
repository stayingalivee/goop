package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func readSource(programPath string) {
    content, err := ioutil.ReadFile(programPath)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(content))
}

