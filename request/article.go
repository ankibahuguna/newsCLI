package request

import (
	"io"
	"net/http"
)

func GetArticle(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
