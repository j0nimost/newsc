package news

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Aljazeera News Headlines
func AljazeeraNews() map[int]string {
	news := make(map[int]string)

	res, err := http.Get("https://www.aljazeera.com/")

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".container .fte-featured__content-wrapper__right .fte-featured__right-inner-articles-wrapper .fte-featured__article-content").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		news[i] = href
		fmt.Printf("Review %d: %s\n", i, title)
	})

	return news
}
