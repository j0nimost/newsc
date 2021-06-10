package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CitizenTv() map[int]string {
	newsLinks := make(map[int]string)
	citizentv := "https://citizentv.co.ke/"
	res, err := http.Get(citizentv)

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	// Load the Html
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".main-story .more-election-stories div").Each(func(i int, s *goquery.Selection) {
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
