package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadArt(filename string) []string {
	var ascii []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		ascii = append(ascii, scan.Text())
	}
	return ascii
}

func Err() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
	os.Exit(1)
}

func Input(inputArry []string) (string, string) {
	var argument, banner string
	args := inputArry
	if len(args) == 1 {
		argument = args[0]
		banner = "standard"
	} else if len(args) == 2 {
		argument = args[0]
		banner = strings.ToLower(args[1])
	} else if len(args) > 2 {
		Err()
	} else if len(args) == 0 {
		Err()
	}
	switch banner {
	case "standard":
		banner = "standard.txt"
	case "shadow":
		banner = "shadow.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"
	default:
		fmt.Println("Unknown banner")
		Err()
	}
	return argument, banner
}

func OutputFile() (string, bool) {
	if len(os.Args) <= 1 {
		Err()
	}
	var output string
	hasFlag := false
	Args := strings.ToLower(os.Args[1])
	if strings.HasPrefix(os.Args[1], "--output=") {
		output = strings.TrimPrefix(os.Args[1], "--output=")
		fileNameCheck := strings.TrimPrefix(Args, "--output=")
		if strings.HasSuffix(fileNameCheck, ".txt") {
			if fileNameCheck != "shadow.txt" && fileNameCheck != "standard.txt" && fileNameCheck != "thinkertoy.txt" {
				hasFlag = true
			}
		} else {
			fmt.Println("File Extension should be txt \nExample abc.txt")
			Err()
		}
	}
	return output, hasFlag
}

func ArgsArry(b bool) []string {
	args := []string{}
	if b {
		args = os.Args[2:]
	} else {
		args = os.Args[1:]
	}
	return args
}

// func Colorize() (bool, string, string, []string) {
// 	checkColorize := false
// 	target := ""
// 	color := ""
// 	inputArry := []string{}
// 	if len(os.Args[1:]) >= 2 && len(os.Args[1:]) <= 4 {
// 		args := os.Args[1:]
// 		if strings.HasPrefix(args[0], "--color=") {
// 			color = strings.TrimPrefix(args[0], "--color=")
// 			checkColorize = true
// 			if len(args) == 4 {
// 				target = args[1]
// 				inputArry = append(inputArry, args[2:]...)
// 				targetCheck(inputArry[0], target)
// 			} else if len(args) == 3 {
// 				asciiArtCheck := strings.ToLower(args[2])
// 				if asciiArtCheck == "shadow" || asciiArtCheck == "standard" || asciiArtCheck == "thinkertoy" {
// 					inputArry = append(inputArry, args[1:]...)
// 				} else {
// 					target = args[1]
// 					inputArry = append(inputArry, args[2:]...)
// 					targetCheck(inputArry[0], target)
// 				}
// 			} else {
// 				inputArry = append(inputArry, args[1:]...)
// 			}
// 		}
// 	} else if len(os.Args[1:]) == 1 {
// 		if strings.HasPrefix(os.Args[1], "--color=") {
// 			Err()
// 		} else {
// 			inputArry = append(inputArry, os.Args[1:]...)
// 		}
// 	} else {
// 		Err()
// 	}

// 	return checkColorize, color, target, inputArry
// }

func Colorize(inputArgs []string) (bool, string, string, []string) {
	checkColorize := false
	target := ""
	color := ""
	inputArry := []string{}
	if len(inputArgs) >= 2 && len(inputArgs) <= 4 {
		args := inputArgs
		if strings.HasPrefix(args[0], "--color=") {
			color = strings.TrimPrefix(args[0], "--color=")
			checkColorize = true
			if len(args) == 4 {
				target = args[1]
				inputArry = append(inputArry, args[2:]...)
				targetCheck(inputArry[0], target)
			} else if len(args) == 3 {
				asciiArtCheck := strings.ToLower(args[2])
				if asciiArtCheck == "shadow" || asciiArtCheck == "standard" || asciiArtCheck == "thinkertoy" {
					inputArry = append(inputArry, args[1:]...)
				} else {
					target = args[1]
					inputArry = append(inputArry, args[2:]...)
					targetCheck(inputArry[0], target)
				}
			} else {
				inputArry = append(inputArry, args[1:]...)
			}
		} else {
			inputArry = inputArgs
		}
	} else if len(inputArgs) == 1 {
		if strings.HasPrefix(inputArgs[0], "--color=") {
			Err()
		} else {
			inputArry = append(inputArry, inputArgs...)
		}
	} else {
		Err()
	}

	return checkColorize, color, target, inputArry
}

func targetCheck(text, target string) {
	if !strings.Contains(text, target) {
		fmt.Println("Check requested letters to be colored")
		Err()
	}
}

func CheckTextInput(text string) {
	if text == "" {
		fmt.Println("No input text")
		Err()
	}
}
