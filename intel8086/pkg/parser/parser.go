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
        // ideally we would make a map with slices of strings and
        // group the values together in this step?
		for k, v := range values {
            var byteSlice []string
            fmt.Printf("current top: %d\n", current)
            fmt.Printf("numbyte top: %d\n", numByte)

            if k != 0 {
                if k <= current {
                    fmt.Println("continuing")
                    continue
                }
            }
            fmt.Println(fmt.Sprintf("%08b", v))

            instruction, numByte, err = differenciateOpcode(fmt.Sprintf("%08b", v), fmt.Sprintf("%08b", values[k+1]))
            if err != nil {
                return err
            }
            fmt.Printf("numbyte after eval: %d\n", numByte)
            _ = instruction
            current = k+(numByte-1)
            end := k

            fmt.Printf("current after eval: %d\n", current)
            fmt.Printf("end after eval: %d\n", current)

            for i := current; i >= end; i-- {
                byteSlice = append(byteSlice, fmt.Sprintf("%08b", values[i]))
            }
            fmt.Println("kajsdf")
            fmt.Println(byteSlice)

            binaries = append(binaries, byteSlice)
            fmt.Println(binaries)

		}

        _ = numByte

		// for i := 1; i < len(binaries); i += 2 {
		// 	err = processInstruction(binaries[i-1], binaries[i],instruction)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

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

// Needs changing, doesn't work anymore
func processInstruction(binary1, binary2 string, instruction string) (error) {
	d := binary1[6]
	w := binary1[7]
	mod := binary2[0:2]
	reg := binary2[2:5]
	rm := binary2[5:]

	if mod != "11" {
		return fmt.Errorf("Unknown mod value")
	}

	sourceReg, destReg, err := registers.IdentifyRegisters(string(d), string(w), reg, rm)
	if err != nil {
		return err
	}

	fmt.Printf("%s %s, %s\n", instruction, destReg, sourceReg)
	return nil
}

