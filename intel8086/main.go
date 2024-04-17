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
    assemblyInstructions := make([]string, 2)

    for scanner.Scan() {
        values := scanner.Bytes()

        for i, v := range values {
            binaries[i] = fmt.Sprintf("%08b", v)
        }

        opcode := binaries[0][0:6]
        d := binaries[0][6]
        w := binaries[0][7]
        fmt.Println(opcode)
        fmt.Println(string(d))
        fmt.Println(string(w))

        mod := binaries[1][0:2] 
        reg := binaries[1][2:5]
        rm := binaries[1][5:]
        fmt.Println(mod)
        fmt.Println(reg)
        fmt.Println(rm)
    }

    fmt.Println(assemblyInstructions)

}
