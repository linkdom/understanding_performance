package test

import (
    "intel8086/pkg/registers"
    "testing"
)

func TestIdentifyRegisters(t *testing.T) {
	tests := []struct {
		name     string
		d        string
		w        string
		reg      string
		rm       string
		wantSrc  string
		wantDest string
		wantErr  bool
	}{
	    {"Valid byte registers", "0", "0", "000", "001", "al", "cl", false},
        {"Valid word registers", "0", "1", "111", "110", "di", "si", false},
        {"Reverse direction", "1", "1", "111", "110", "si", "di", false},
        {"Invalid d", "2", "0", "000", "001", "", "", true},
        {"Invalid w", "0", "2", "000", "001", "", "", true},
        {"Invalid reg", "0", "0", "010", "999", "", "", true},
        {"Invalid rm", "0", "0", "999", "001", "", "", true},}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, dest, err := registers.IdentifyRegisters(tt.d, tt.w, tt.reg, tt.rm)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentifyRegisters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (src != tt.wantSrc || dest != tt.wantDest) {
				t.Errorf("IdentifyRegisters() gotSrc = %v, gotDest = %v, wantSrc = %v, wantDest = %v", src, dest, tt.wantSrc, tt.wantDest)
			}
		})
	}
}

