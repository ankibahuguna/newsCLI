package parser

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(html io.Reader) (string, error) {
	var e error
	body := ""

	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		return "", err
	}
	doc.Find("div.author-bottom").Parent().Prev().Children().Each(func(i int, s *goquery.Selection) {

		idString, _ := s.Attr("id")

		if strings.Contains(idString, "content-body-") {
			s.Children().Each(func(j int, el *goquery.Selection) {
				body = body + el.Text()
				body = body + "\n\n"
			})

			return
		}
	})
	return body, e
}
