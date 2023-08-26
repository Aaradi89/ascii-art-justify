package pkg

import (
	"fmt"
	"os"
	"strings"
)

func Output(fileName string, asciiArt []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	for i, line := range asciiArt {
		if i < len(asciiArt) {
			fmt.Fprintln(file, line)
		}
	}
}

func OutputT(outputFile, text, align string, asciiArt []string, toTxt, alignCheck bool) {
	finalArt := make([]string, 0)
	if IsNewLine(text) {
		args := strings.Split(text, "\\n")
		args = FixNewLines(args)
		for _, arg := range args {
			if arg == "" {
				finalArt = append(finalArt, "")
				continue
			}
			if alignCheck {
				sizeWithSpace, sizeWithoutSpace := ArtSize(arg, asciiArt)
				arg = TextAlign(align, arg, sizeWithSpace, sizeWithoutSpace)
			}
			addArt := ArtPreparation(arg, asciiArt, alignCheck)
			fmt.Println(alignCheck)
			finalArt = append(finalArt, addArt...)
		}
	} else if text != "" {
		if alignCheck {
			sizeWithSpace, sizeWithoutSpace := ArtSize(text, asciiArt)
			text = TextAlign(align, text, sizeWithSpace, sizeWithoutSpace)
		}
		finalArt = ArtPreparation(text, asciiArt, alignCheck)
	}
	if toTxt {
		Output(outputFile, finalArt)
	} else {
		PrintArt(finalArt)
	}
}
