package main

import (
    "testing"
)

func TestIdentifyRegisters(t *testing.T) {

    tests := []struct {
        name    string
        d       string
        w       string
        reg     string
        rm      string
        wantSrc string
        wantDst string
        wantErr bool
    }{
        {"Valid byte registers", "0", "0", "000", "001", "al", "cl", false},
        {"Valid word registers", "0", "1", "111", "110", "di", "si", false},
        {"Reverse direction", "1", "1", "111", "110", "si", "di", false},
        {"Invalid d", "2", "0", "000", "001", "", "", true},
        {"Invalid w", "0", "2", "000", "001", "", "", true},
        {"Invalid reg", "0", "0", "010", "999", "", "", true},
        {"Invalid rm", "0", "0", "999", "001", "", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotSrc, gotDst, err := identifyRegisters(tt.d, tt.w, tt.reg, tt.rm)
            if (err != nil) != tt.wantErr {
                t.Errorf("identifyRegisters() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if gotSrc != tt.wantSrc || gotDst != tt.wantDst {
                t.Errorf("identifyRegisters() gotSrc = %v, gotDst = %v, wantSrc = %v, wantDst = %v", gotSrc, gotDst, tt.wantSrc, tt.wantDst)
            }
        })
    }
}

