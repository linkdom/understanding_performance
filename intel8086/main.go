package main

import (
	"flag"
	"log"
	"path/filepath"

	"intel8086/pkg/parser"
	"intel8086/pkg/utils"
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

    outputFile := utils.OutputFileName(*input)

    if err := parser.ProcessFile(path, outputFile); err != nil {
        log.Fatal(err)
    }
}

