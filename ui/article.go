package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
)

func RenderArticle(text string) {

	tm.Clear()
	d := color.New(color.FgWhite, color.Italic)
	padded := d.Sprintf("%-72v", text)
	pager := "/usr/bin/more"
	if runtime.GOOS == "windows" {
		pager = "C:\\Windows\\System32\\more.com"
	}
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(padded)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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
