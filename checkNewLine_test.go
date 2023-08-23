package main

import (
	"ascii-art/pkg"
	"testing"
)

func testIsNewline(t *testing.T) {
	if !pkg.IsNewLine("hello \\n world") {
		t.Error("Newline should be true")
	}
}
