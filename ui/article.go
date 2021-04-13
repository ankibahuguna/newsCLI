package ui

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tm "github.com/buger/goterm"
)

func RenderArticle(text string) {

	tm.Clear()
	//d := color.New(color.FgWhite, color.Italic)
	//padded := d.Sprintf("%-72v", text)
	pager := "/usr/bin/bat"
	if runtime.GOOS == "windows" {
		pager = "C:\\Windows\\System32\\more.com"
	}
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(text)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
