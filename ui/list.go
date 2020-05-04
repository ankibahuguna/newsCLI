package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ShowHeadLines(label string, size int, items []string) (int, error) {

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
