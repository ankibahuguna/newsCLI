package ui

import (
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
}
