package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    file, err := os.Open("/home/dom/development/go/understanding_performance/intel8086/single_register_move")
    if err != nil {
         log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    binaries := make([]string, 2)

    for scanner.Scan() {
        values := scanner.Bytes()

        for i, v := range values {
            binaries[i] = fmt.Sprintf("%08b", v)
        }
    }

    for _, x := range binaries {
        fmt.Println(x)
    }




}
