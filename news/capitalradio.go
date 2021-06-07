package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Capital Radio News Listing
func Capitalradio() map[int]string {
	news := make(map[int]string)

	res, err := http.Get("https://www.capitalfm.co.ke/news/")

	if err != nil {
		log.Fatalf("Status Code Error: %d : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".zox-feat-right-wrap .zox-side-list-wrap section").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title // .marquee .marquee-content .items
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		news[i] = href
		trimed := strings.TrimSpace(title)
		fmt.Printf("Review %d: %s\n", i, trimed)
	})

	return news
}
