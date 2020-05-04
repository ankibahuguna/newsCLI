package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tm "github.com/buger/goterm"
)

func ShowPrompt() {

	var input string
	fmt.Scanln(&input)

	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Press `q` to quit any other key to continue >> ")
	ans, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.TrimSpace(ans) == "q" {
		fmt.Println("Bye")
		os.Exit(0)
	}
	tm.Clear()
}
