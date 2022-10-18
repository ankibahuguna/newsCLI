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

	f := Formatter{
		Writer: os.Stdout,
		Indent: []byte(" "),
		Width:  120,
	}
	tm.Clear()
	//	_, err := f.Write([]byte(padded))
	//d := color.New(color.FgWhite, color.Italic)
	pager := "/usr/bin/less"
	if runtime.GOOS == "windows" {
		pager = "C:\\Windows\\System32\\more.com"
	}
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(string(f.format([]byte(text))))
	cmd.Stdout = os.Stdout
	color.Set(color.FgYellow)
	defer color.Unset()
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
