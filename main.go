package main

import (
	"ascii-art/pkg"
	"fmt"
)

func main() {
	alignCheck, align, inputArgs := pkg.AlignFlagCheck()
	asciiColorCheck, color, target, inputArry := pkg.Colorize(inputArgs)
	outputFile, fileNameChecker := pkg.OutputFile()
	//var sizeWithSpace, sizeWithoutSpace int
	if !asciiColorCheck && !alignCheck {
		inputArry = pkg.ArgsArry(fileNameChecker)
	}
	text, banner := pkg.Input(inputArry)
	pkg.CheckTextInput(text)
	asciiArt := pkg.ReadArt(banner)
	//test justify
	// if alignCheck {
	// 	sizeWithSpace, sizeWithoutSpace = pkg.ArtSize(text, asciiArt)
	// 	// 	text = pkg.TextAlign(align, text, sizeWithSpace, sizeWithoutSpace)
	// }
	//fmt.Println(txt)
	//test justify
	if asciiColorCheck {
		color = pkg.ColorCodeCheck(color)
		pkg.ColorOutput(text, color, target, align, asciiArt, alignCheck)
	} else {
		pkg.OutputT(outputFile, text, align, asciiArt, fileNameChecker, alignCheck)
	}
	lenth(167)
}

func lenth(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("-")
	}
}
