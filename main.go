package main

import "ascii-art/pkg"

func main() {
	pkg.GetTerminalWidth()
	asciiColorCheck, color, target, inputArry := pkg.Colorize()
	outputFile, fileNameChecker := pkg.OutputFile()
	if !asciiColorCheck {
		inputArry = pkg.ArgsArry(fileNameChecker)
	}
	text, banner := pkg.Input(inputArry)
	pkg.CheckTextInput(text)
	asciiArt := pkg.ReadArt(banner)
	if asciiColorCheck {
		color = pkg.ColorCodeCheck(color)
		pkg.ColorOutput(text, color, target, asciiArt)
	} else {
		pkg.OutputT(outputFile, text, asciiArt, fileNameChecker)
	}
}
