package registers

import "errors"

func IdentifyRegisters(d string, w string, reg string, rm string) (sourceReg string, destReg string, err error) {
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
