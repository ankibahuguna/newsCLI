package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	tm "github.com/buger/goterm"
	"github.com/eidolon/wordwrap"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mmcdole/gofeed"
)

type News struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

var news []News
var headlines []string

func main() {

	args := os.Args[1:]

	category := "feeder/default.rss"

	if len(args) > 0 {
		category = strings.TrimSpace(args[0]) + "/" + category
	}

	url := fmt.Sprintf("https://www.thehindu.com/%v", category)

	var err error
	news, err = getNews(url)

	if err != nil {
		fmt.Println(err)
		panic("Something went wrong")
	}

	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	for _, val := range news {
		title, description := val.Title, val.Description
		totalLength := len(title+description) - len(description)

		description = description[0:min(len(description), totalLength)] + "..."
		titleString := fmt.Sprintf("%s (%s)", white(title), green(description))
		headlines = append(headlines, titleString)
	}

	renderOutPut()

}

func renderOutPut() {

	index, err := promptUI("HeadLines", 20, headlines)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	content, err := parseHTML(news[index].Link)

	if err != nil {
		fmt.Println("Some shit went wrong", err)
	}
	outPutToTerminal(content)
}

func promptUI(label string, size int, items []string) (int, error) {

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

func getNews(url string) ([]News, error) {

	var news []News

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)

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
		newsItem := News{Title: title, Description: description, Link: link}
		news = append(news, newsItem)
	}

	return news, nil
}

func parseHTML(url string) (string, error) {
	body, err := getArticle(url)
	if err != nil {
		return "", err
	}
	html, err := getText(body)

	if err != nil {
		return "", err
	}
	return html, nil
}

func getArticle(url string) (io.Reader, error) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(url))

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func getText(html io.Reader) (string, error) {
	var e error
	body := ""

	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		return "", err
	}
	wrapper := wordwrap.Wrapper(120, false)
	doc.Find("div.author-bottom").Parent().Prev().Children().Each(func(i int, s *goquery.Selection) {

		idString, _ := s.Attr("id")

		if strings.Contains(idString, "content-body-") {
			s.Children().Each(func(j int, el *goquery.Selection) {
				body = body + wrapper(el.Text())
				body = body + "\n\n"
			})

			return
		}
	})

	return wrapper(body), e
}

func outPutToTerminal(text string) {
	tm.Clear()
	d := color.New(color.FgWhite, color.Italic)
	padded := d.Sprintf("%-72v", text)
	pager := "/usr/bin/more"
	if runtime.GOOS == "windows" {
		pager = "C:\\Windows\\System32\\more.com"
	}
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(padded)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	var input string
	fmt.Scanln(&input)

	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Press `q` to quit any other key to continue >> ")
	ans, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.TrimSpace(ans) == "q" {
		fmt.Println("Bye")
		os.Exit(0)
	}
	tm.Clear()
	renderOutPut()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
