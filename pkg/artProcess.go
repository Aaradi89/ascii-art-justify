package pkg

import (
	"fmt"
	"os"
)

func ArtPreparation(text string, art []string) []string {
	result := make([]string, 8)
	for _, r := range text {
		if r >= ' ' && r <= '~' {
			position := ((int(r) - ' ') * 9) + 1
			for i := 0; i < 8; i++ {
				result[i] += art[position+i]
			}

		} else {
			fmt.Println("Invalid input")
			Err()
			os.Exit(1)
		}
	}
	return result
}

func PrintArt(art []string) {
	for _, r := range art {
		fmt.Println(r)
	}
}

func FixNewLines(text []string) []string {
last := len(text) - 1
if text[last] == "" && last > 0 {
    return text[:last]
}
return text
}
