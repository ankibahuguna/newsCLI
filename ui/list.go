package ui

import (
	"fmt"

	"github.com/ankibahuguna/news/types"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ShowHeadLines(label string, size int, news []types.News) (int, error) {

	items := getHeadLines(news)

	color.Set(color.FgCyan)
	defer color.Unset()
	prompt := promptui.Select{
		Label: label,
		Items: items,
		Size:  size,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}

	return index, nil
}

func getHeadLines(news []types.News) []string {
	var headlines []string

	red := color.New(color.Bold, color.FgWhite).SprintFunc()

	for _, val := range news {
		title := val.Title
		title = fmt.Sprintf(red(title))
		headlines = append(headlines, title)
	}
	return headlines
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
