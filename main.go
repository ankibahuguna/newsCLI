package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ankibahuguna/news/parser"
	"github.com/ankibahuguna/news/request"
	"github.com/ankibahuguna/news/types"
	"github.com/ankibahuguna/news/ui"
	"github.com/fatih/color"
)

func main() {

	args := os.Args[1:]

	category := "feeder/default.rss"

	if len(args) > 0 {
		category = strings.TrimSpace(args[0]) + "/" + category
	}

	url := fmt.Sprintf("https://www.thehindu.com/%v", category)

	news, err := request.GetNews(url)

	if err != nil {
		fmt.Println(err)
		panic("Something went wrong")
	}

	headlines := getHeadLines(news)
	index, err := ui.ShowPrompt("HeadLines", 20, headlines)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}

	body, err := request.GetArticle(news[index].Link)

	if err != nil {
		fmt.Println("Some shit went wrong", err)
		panic(err)
	}

	content, err := parser.ParseHTML(body)

	if err != nil {
		panic(err)
	}

	ui.RenderArticle(content)

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
