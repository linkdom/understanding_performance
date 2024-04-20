package main

import (
	"flag"
	"log"
	"path/filepath"

	"intel8086/pkg/parser"
)

func main() {
    input := flag.String("input", "", "path including filename")
    flag.Parse()

    if *input == "" {
        log.Fatal("No path provided")
    }

    path, err := filepath.Abs(*input)
    if err != nil {
        log.Fatal(err)
    }

    if err := parser.ProcessFile(path); err != nil {
        log.Fatal(err)
    }
}

