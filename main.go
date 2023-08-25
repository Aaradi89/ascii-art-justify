package main

import (
	"ascii-art/pkg"
	"fmt"
)

func main() {
	asciiColorCheck, color, target, inputArry := pkg.Colorize()
	outputFile, fileNameChecker := pkg.OutputFile()
	if !asciiColorCheck {
		inputArry = pkg.ArgsArry(fileNameChecker)
	}
	text, banner := pkg.Input(inputArry)
	pkg.CheckTextInput(text)
	asciiArt := pkg.ReadArt(banner)
	//test justify
	sizeWithSpace, sizeWithoutSpace := pkg.ArtSize(text, asciiArt)
	text = pkg.TextAlign("right", text, sizeWithSpace, sizeWithoutSpace)
	//fmt.Println(txt)
	//test justify
	if asciiColorCheck {
		color = pkg.ColorCodeCheck(color)
		pkg.ColorOutput(text, color, target, asciiArt)
	} else {
		pkg.OutputT(outputFile, text, asciiArt, fileNameChecker)
	}
	lenth(179)
}

func lenth(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("-")
	}
}
