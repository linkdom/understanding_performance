package utils

import "path/filepath"

// OutputFileName generates the output file name based on the input file.
func OutputFileName(input string) string {
	return filepath.Dir(input) + "/output_" + filepath.Base(input) + ".asm"
}

