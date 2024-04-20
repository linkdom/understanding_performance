package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"intel8086/pkg/registers"
)

func ProcessFile(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(file)
	var binaries []string

	_, err = f.WriteString("bits 16\n\n")
	if err != nil {
		return err
	}

	for scanner.Scan() {
		values := scanner.Bytes()

		for _, v := range values {
			binaries = append(binaries, fmt.Sprintf("%08b", v))
		}

        numByte, err := differenciateOpcode(binaries[0])
        if err != nil {
            return err
        }
        _ = numByte

		for i := 1; i < len(binaries); i += 2 {
			err = processInstruction(binaries[i-1], binaries[i], f)
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
func differenciateOpcode(binary string) (uint8, error) {

    if binary[0:4] == "1011" {
        if string(binary[5]) == "1" {
            return 3, nil
        }
        return 2, nil
    }

    if binary[0:6] == "100010" {
        return 2, nil
    }

    switch binary[0:7] {
    case "1100011":
        if string(binary[8]) == "1" {
            return 6, nil
        }
        return 5, nil
    case "1010000", "1010001":
        if string(binary[8]) == "1" {
            return 3, nil
        }
        return 2, nil
    default:
        return 0, errors.New("Unknown Opcode")
    }

}

func processInstruction(binary1, binary2 string, f *os.File) (error) {
	opcode := binary1[0:6]
	d := binary1[6]
	w := binary1[7]
	mod := binary2[0:2]
	reg := binary2[2:5]
	rm := binary2[5:]

    var instruction string

	if opcode == "100010" {
		instruction = "mov"
	}

	if mod != "11" {
		return fmt.Errorf("Unknown mod value")
	}

	sourceReg, destReg, err := registers.IdentifyRegisters(string(d), string(w), reg, rm)
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintf("%s %s, %s\n", instruction, destReg, sourceReg))
	if err != nil {
		return err
	}

	return nil
}

