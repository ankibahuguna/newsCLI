package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowPrompt() (bool, error) {

	var input string
	fmt.Scanln(&input)

	buf := bufio.NewReader(os.Stdin)
	ans, err := buf.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if strings.TrimSpace(ans) == "q" {
		return false, nil
	}

	return true, nil
}
