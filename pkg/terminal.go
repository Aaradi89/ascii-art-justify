package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out = out[:len(out)-1]
	col, err := strconv.Atoi(string(out))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return col
}

func ArtSize(text string, asciiArt []string) (int, int) {
	sizeWithSpace := 0
	sizeWithoutSpace := 0
	for _, r := range text {
		artIndex := (r-' ')*9 + 1
		if r == ' ' {
			sizeWithoutSpace += 0
		} else {
			sizeWithoutSpace += len(asciiArt[artIndex])
		}
		sizeWithSpace += len(asciiArt[artIndex])
	}
	return sizeWithSpace, sizeWithoutSpace
}

func TextAlign(Align, Text string, sizeWithSpace, sizeWithoutSpace int) string {
	align := Align
	text := Text
	terminalWidth := GetTerminalWidth()
	rightAlignCols := terminalWidth - sizeWithSpace
	centerAlignCols := (terminalWidth - sizeWithSpace) / 2
	justifyAlignCols := terminalWidth - sizeWithoutSpace
	fmt.Printf("term : %v; right : %v; center : %v \n", terminalWidth, rightAlignCols, centerAlignCols)
	fmt.Printf("with space : %v; without space : %v\n", sizeWithSpace, sizeWithoutSpace)
	spaces := ""

	switch align {
	case "right":
		for i := 0; i < rightAlignCols; i++ {
			spaces += " "
		}
		text = strings.ReplaceAll(text, " ", "      ")
		text = spaces + text
	case "center":
		for i := 0; i < centerAlignCols; i++ {
			spaces += " "
		}
		text = strings.ReplaceAll(text, " ", "      ")
		text = spaces + text + spaces
	case "justify":
		textArry := make([]string, 0)
		reg1 := regexp.MustCompile(`(.+)(\s+)$`)
		text = reg1.ReplaceAllString(text, "$1")
		reg2 := regexp.MustCompile(`^(\s+)(.+)`)
		text = reg2.ReplaceAllString(text, "$2")
		fmt.Println(text)
		textArry0 := strings.Split(text, " ")
		for _, ele := range textArry0 {
			if ele != "" {
				textArry = append(textArry, ele)
			}
		}
		justifyAlignCols /= (len(textArry) - 1)
		for i := 0; i < justifyAlignCols; i++ {
			spaces += " "
		}
		text = ""
		for i := 0; i < len(textArry); i++ {
			text += textArry[i]
			if i < len(textArry)-1 {
				text += spaces
			}
		}
	default:
		text = strings.ReplaceAll(text, " ", "      ")
		return text
	}
	return text
}
