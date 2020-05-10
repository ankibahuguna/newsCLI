package request

import (
	"errors"
	"log"
	"strings"

	"github.com/ankibahuguna/news/types"
	"github.com/mmcdole/gofeed"
)

func GetArticleList(feedUrl string) ([]types.News, error) {

	var news []types.News

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUrl)

	if err != nil {
		log.Println(err, "Some shit went wrong")
		return nil, errors.New("Couldn't parse RSS feed")
	}

	for _, item := range feed.Items {
		var (
			title       = strings.TrimSpace(item.Title)
			description = strings.TrimSpace(item.Description)
			link        = strings.TrimSpace(item.Link)
		)
		newsItem := types.News{Title: title, Description: description, Link: link}
		news = append(news, newsItem)
	}

	return news, nil
}
