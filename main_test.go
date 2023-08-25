package main

import (
	"ascii-art/pkg"
	"testing"
)

func TestArtPreperation(t *testing.T) {
	str := `PASS`
	asciiArt := pkg.ReadArt("standard.txt")
	art := pkg.ArtPreparation(str, asciiArt, false)
	pkg.PrintArt(art)
}
