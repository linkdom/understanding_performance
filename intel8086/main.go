package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
    file, err := os.Open("/home/dom/development/go/understanding_performance/intel8086/many_register_move")
    if err != nil {
         log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var binaries []string

    fmt.Println("bits 16\n")

    for scanner.Scan() {
        values := scanner.Bytes()
        var instruction string

        for _, v := range values {
            binaries = append(binaries, fmt.Sprintf("%08b", v))
        }

        for i := 1; i < len(binaries); i+=2 {
            opcode := binaries[i-1][0:6]
            d := binaries[i-1][6]
            w := binaries[i-1][7]
            mod := binaries[i][0:2] 
            reg := binaries[i][2:5]
            rm := binaries[i][5:]
            _ = mod

            if opcode == "100010" {
                instruction = "mov"
            }

            sourceReg, destReg, err := identifyRegisters(string(d), string(w), reg, rm)
            if err != nil {
                fmt.Println(err)
            }

            fmt.Printf("%s %s, %s\n", instruction, destReg, sourceReg)

        }

    }
}

func identifyRegisters(d string, w string, reg string, rm string) (sourceReg string, destReg string, err error) {
	registers := map[string]map[string]string{
		"0": {
			"000": "al",
			"001": "cl",
			"010": "dl",
			"011": "bl",
			"100": "ah",
			"101": "ch",
			"110": "dh",
			"111": "bh",
		},
		"1": {
			"000": "ax",
			"001": "cx",
			"010": "dx",
			"011": "bx",
			"100": "sp",
			"101": "bp",
			"110": "si",
			"111": "di",
		},
	}

	table, ok := registers[w]
	if !ok {
		return "", "", errors.New("Invalid 'w' value")
	}

	sourceReg, ok = table[reg]
	if !ok {
		return "", "", errors.New("Invalid 'reg' value")
	}

	destReg, ok = table[rm]
	if !ok {
		return "", "", errors.New("Invalid 'rm' value")
	}

	if d == "1" {
		sourceReg, destReg = destReg, sourceReg
	} else if d != "0" {
		return "", "", errors.New("Invalid 'd' value")
	}

	return sourceReg, destReg, nil
}
