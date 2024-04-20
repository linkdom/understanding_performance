package test

import (
    "intel8086/pkg/parser"
    "path/filepath"
    "testing"
)

func TestProcessFile(t *testing.T) {
	inputPath := filepath.Join("/home/dom/development/learning/understanding_performance/intel8086/files/many_register_move")

	err := parser.ProcessFile(inputPath)
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	// expected := `bits 16

// mov cx, bx
// mov ch, ah
// mov dx, bx
// mov si, bx
// mov bx, di
// mov al, cl
// mov ch, ch
// mov bx, ax
// mov bx, si
// mov sp, di
// mov bp, ax
// `

	// if string(content) != expected {
	// 	t.Errorf("Expected %v, got %v", expected, string(content))
	// }
}

