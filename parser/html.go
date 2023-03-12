package parser

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func FilterOutCommentShareWidget(i int, s *goquery.Selection) bool {
	val, exists := s.Parent().Attr("class")

	if exists == true && val == "comments-shares share-page" {
		return false
	}

	return true
}

func ParseHTML(html io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		return "", err
	}

	body := ""

	doc.Find("div.articlebodycontent").Find("p").Not(".related-topics-list").FilterFunction(FilterOutCommentShareWidget).Each(func(j int, el *goquery.Selection) {
		body = body + el.Text()
		body = body + "\n\n"
	})

	return body, nil
}
