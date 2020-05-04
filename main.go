package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ankibahuguna/news/parser"
	"github.com/ankibahuguna/news/request"
	"github.com/ankibahuguna/news/ui"
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

	for {
		index, err := ui.ShowHeadLines("HeadLines", 20, news)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			panic(err)
		}

		body, err := request.GetArticle(news[index].Link)

		if err != nil {
			panic(err)
		}

		content, err := parser.ParseHTML(body)

		if err != nil {
			panic(err)
		}

		ui.RenderArticle(content)
		ui.ShowPrompt()
	}

}
