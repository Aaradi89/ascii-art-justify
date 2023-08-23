package pkg

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// to save the letters and there colors in one object
type colorCode struct {
	letter rune
	color  string
}

// to save the letters and there requird colors in one object and return them in one arry
func TCreateColorCode(w, c, t string) []colorCode {

	var colorcodes []colorCode
	var letter colorCode
	word := w
	color := c
	target := t
	reset := "\033[0m"

	for i := 0; i < len(word); i++ {
		if len(word[i:]) >= len(target) {
			targetLen := i + len(target)
			if word[i:targetLen] == target {
				for i < targetLen {
					letter.letter = rune(word[i])
					letter.color = color
					colorcodes = append(colorcodes, letter)
					i++
				}
			}
		}
		if i < len(word) {
			letter.letter = rune(word[i])
			if t == "" {
				letter.color = color
			} else {
				letter.color = reset
			}
			colorcodes = append(colorcodes, letter)
		}
	}
	return colorcodes
}

// take the return from createColorcodes and the return from readArt to create a 8 arry of colored content ready to be printed
func ColorArtPreparation(text []colorCode, art []string) []string {
	result := make([]string, 8)
	for _, r := range text {
		if r.letter >= ' ' && r.letter <= '~' {
			position := ((int(r.letter) - ' ') * 9) + 1
			for i := 0; i < 8; i++ {
				result[i] += r.color + art[position+i]
			}

		} else {
			fmt.Println("Invalid input")
			Err()
			os.Exit(1)
		}
	}
	return result
}

// check the validity of the requested color
func ColorCodeCheck(colorCode string) string {
	colorCode = strings.ToLower(colorCode)
	red := "\033[38;2;255;0;0m"
	green := "\033[38;2;0;255;0m"
	blue := "\033[38;2;0;0;255m"
	yellow := "\033[38;2;255;255;0m"
	cyan := "\033[38;2;0;255;255m"
	purple := "\033[38;2;255;0;255m"
	orange := "\033[38;2;255;125;0m"
	black := "\033[38;2;0;0;0m"
	white := "\033[38;2;255;255;255m"
	var color string

	switch colorCode {
	case "red":
		color = red
	case "green":
		color = green
	case "blue":
		color = blue
	case "yellow":
		color = yellow
	case "cyan":
		color = cyan
	case "purple":
		color = purple
	case "orange":
		color = orange
	case "black":
		color = black
	case "white":
		color = white
	default:
		reg := regexp.MustCompile(`^rgb\((\d+),(\d+),(\d+)\)$`)
		if reg.MatchString(colorCode) {
			redCode := reg.ReplaceAllString(colorCode, `$1`)
			greenCode := reg.ReplaceAllString(colorCode, `$2`)
			blueCode := reg.ReplaceAllString(colorCode, `$3`)
			redTest, _ := strconv.ParseInt(redCode, 10, 64)
			greenTest, _ := strconv.ParseInt(greenCode, 10, 64)
			blueTest, _ := strconv.ParseInt(blueCode, 10, 64)
			if redTest >= 0 && greenTest >= 0 && blueTest >= 0 && redTest <= 255 && greenTest <= 255 && blueTest <= 255 {
				color = "\033[38;2;" + redCode + ";" + greenCode + ";" + blueCode + "m"
			} else {
				colorErr()
			}
		} else {
			colorErr()
		}
	}
	return color

}

// use when there is N ERROR IN COLOR CODES
func colorErr() {
	fmt.Println("wrong color code")
	fmt.Println("Use the Following Format")
	fmt.Println("rgb(0-255,0-255,0-255)")
	fmt.Println("Example : rgb(25,155,255)")
	os.Exit(1)
}

func ColorOutput(text, color, target string, asciiArt []string) {
	finalArt := make([]string, 0)
	if IsNewLine(text) {
		args := strings.Split(text, "\\n")
		args = FixNewLines(args)
		for _, arg := range args {
			if arg == "" {
				finalArt = append(finalArt, "")
				continue
			}
			colorArry := CreateColorCode(arg, color, target)
			addArt := ColorArtPreparation(colorArry, asciiArt)
			finalArt = append(finalArt, addArt...)
		}
	} else if text != "" {
		colorArry := CreateColorCode(text, color, target)
		finalArt = ColorArtPreparation(colorArry, asciiArt)
	}
	PrintArt(finalArt)
	fmt.Printf("\033[0m")
}

func CreateColorCode(w, c, t string) []colorCode {

	var colorcodes []colorCode
	var letter colorCode
	word := w
	color := c
	target := t
	reset := "\033[0m"
	for i := 0; i < len(word); i++ {
		if len(word[i:]) >= len(target) {
			targetLen := i + len(target)
			if word[i:targetLen] == target {
				for i < targetLen {
					letter.letter = rune(word[i])
					letter.color = color
					colorcodes = append(colorcodes, letter)
					i++
				}
				if target != "" {
					i--
					continue
				}
			}
		}
		if i < len(word) {
			letter.letter = rune(word[i])
			if t == "" {
				letter.color = color
			} else {
				letter.color = reset
			}
			colorcodes = append(colorcodes, letter)
		}
	}
	return colorcodes
}
