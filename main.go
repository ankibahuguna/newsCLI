package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ankibahuguna/news/parser"
	"github.com/ankibahuguna/news/request"
	"github.com/ankibahuguna/news/ui"
	tm "github.com/buger/goterm"
)

func main() {

	args := os.Args[1:]

	category := "feeder/default.rss"

	if len(args) > 0 {
		category = strings.TrimSpace(args[0]) + "/" + category
	}

	feedUrl := fmt.Sprintf("https://www.thehindu.com/%v", category)

	news, err := request.GetArticleList(feedUrl)

	if err != nil {
		fmt.Println(err)
		panic("Could not parse RSS feed.")
	}

	for {
		index, err := ui.ShowHeadLines("HeadLines", 20, news)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			panic(err)
		}

		body, err := request.GetArticle(news[index].Link)

		if err != nil {
			fmt.Println(err)
			panic("Could not get the article")
		}

		content, err := parser.ParseHTML(body)

		if err != nil {
			panic(err)
		}

		ui.RenderArticle(content)
		fmt.Println()
		fmt.Println("Press `q` to quit any other key to continue >> ")
		ans, _ := ui.ShowPrompt()

		if !ans {
			fmt.Println("Exiting")
			break
		}
		tm.Clear()

	}

}
