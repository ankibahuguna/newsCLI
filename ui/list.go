package ui

import (
	"fmt"

	"github.com/ankibahuguna/news/types"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ShowHeadLines(label string, size int, news []types.News) (int, error) {

	items := getHeadLines(news)

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
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	for _, val := range news {
		title, description := val.Title, val.Description
		totalLength := len(title+description) - len(description)

		description = description[0:min(len(description), totalLength)] + "..."
		titleString := fmt.Sprintf("%s (%s)", white(title), green(description))
		headlines = append(headlines, titleString)
	}

	return headlines
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
