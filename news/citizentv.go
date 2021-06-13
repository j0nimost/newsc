package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CitizenTv struct {
	Url, Query string
}

func (ctv CitizenTv) GetNews() map[int] string {
	newsLinks := make(map[int]string)
	res, err := http.Get(ctv.Url)

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	// Load the Html
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(ctv.Query).Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")

		if href != "#" {
			newsLinks[i] = href
			trim := strings.TrimSpace(title)
			fmt.Printf("Review %d: %s\n", i, trim)
		}

	})

	return newsLinks
}