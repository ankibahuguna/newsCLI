package parser

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(html io.Reader) (string, error) {
	var e error
	body := ""

	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	doc.Find("div.article").Children().Each(func(i int, s *goquery.Selection) {

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
