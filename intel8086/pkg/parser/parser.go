package parser

import (
	"bufio"
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

		for i := 1; i < len(binaries); i += 2 {
			err = processInstruction(binaries[i-1], binaries[i], f)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

