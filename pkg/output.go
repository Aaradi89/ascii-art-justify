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


func OutputT(outputFile, text string, asciiArt []string, toTxt, alignCheck bool) {
	finalArt := make([]string, 0)
	if IsNewLine(text) {
		args := strings.Split(text, "\\n")
		args = FixNewLines(args)
		for _, arg := range args {
			if arg == "" {
				finalArt = append(finalArt, "")
				continue
			}
			addArt := ArtPreparation(arg, asciiArt,alignCheck)
			finalArt = append(finalArt, addArt...)
		}
	} else if text != "" {
		finalArt = ArtPreparation(text, asciiArt,alignCheck)
	}
	if toTxt {
		Output(outputFile, finalArt)
	} else {
		PrintArt(finalArt)
	}
}
