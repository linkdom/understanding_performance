package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"intel8086/pkg/registers"
)

func ProcessFile(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var binaries [][]string

	fmt.Printf("bits 16\n\n")

	for scanner.Scan() {
		values := scanner.Bytes()
        numByte := 0
        current := 0
        var instruction string

        fmt.Println(fmt.Sprintf("%08b", values))

        // There is an issue with a mov [bp, si], cl instruction
        // I don't yet understand why, i should get two bytes in this
        // instruction and currently i only receive one byte (10001000)
        // this needs further debugging but for now i have commented this
        // assembly instruction to continue working
		for k, v := range values {
            var byteSlice []string

            if k != 0 {
                if k <= current {
                    fmt.Println("continuing")
                    continue
                }
            }

            instruction, numByte, err = differenciateOpcode(fmt.Sprintf("%08b", v), fmt.Sprintf("%08b", values[k+1]))
            if err != nil {
                return err
            }

            current = k+(numByte-1)
            end := k

            for i := current; i >= end; i-- {
                byteSlice = append(byteSlice, fmt.Sprintf("%08b", values[i]))
            }

            binaries = append(binaries, byteSlice)

		}

        for _, v := range binaries {
			err = processInstruction(v, instruction)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// In here i need to figure out what opcode we have got
// so i will know how many bytes i need to process after this 
// byte, the int tells me how many bytes this instruction has
func differenciateOpcode(binary, secondaryByte string) (string, int, error) {
    fmt.Println(binary, secondaryByte)
    fmt.Println(secondaryByte[0:2])

    if binary[0:4] == "1011" {
        if string(binary[4]) == "1" {
            return "mov", 3, nil
        }
        return "mov", 2, nil
    }

    if binary[0:6] == "100010" {

        if secondaryByte[0:2] == "00" {
            return "mov", 2, nil
        } else if secondaryByte[0:2] == "01" {
            return "mov", 3, nil
        } else if secondaryByte[0:2] == "10" {
            return "mov", 4, nil
        }


        return "mov", 2, nil
    }

    switch binary[0:7] {
    case "1100011":
        if string(binary[8]) == "1" {
            return "mov", 6, nil
        }
        return "mov", 5, nil
    case "1010000", "1010001":
        if string(binary[8]) == "1" {
            return "mov", 3, nil
        }
        return "mov", 2, nil
    default:
        return "", 0, errors.New("Unknown Opcode")
    }

}

// the slice is backwards so i need to loop that way
func processInstruction(binaries []string, instruction string) (error) {
    var d string
    var w string
    var mod string
    var reg string
    var rm string

    //the first byte that comes is the one with the opcode
    //i need to check how i then proceed with the following bytes
    //because they have variable length
    for i := len(binaries); i >= 0; i-- {
        d = binaries[i][6:6]
        w = binaries[i][7:7]
        mod = binaries[i][0:2]
        reg = binaries[i][2:5]
        rm = binaries[i][5:]
    }
    _ = mod


	sourceReg, destReg, err := registers.IdentifyRegisters(string(d), string(w), reg, rm)
	if err != nil {
		return err
	}

	fmt.Printf("%s %s, %s\n", instruction, destReg, sourceReg)
	return nil
}

