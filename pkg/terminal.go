package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	GetTerminalWidth()
}

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
