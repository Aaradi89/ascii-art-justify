package main

import (
	"ascii-art/pkg"
	"testing"
)

func TestReadArt(t *testing.T) {
	lenTest := len(pkg.ReadArt(`standard.txt`))
	if lenTest != 855 {
		t.Errorf(`got %d, want %d`, lenTest, 855)
	}
}
