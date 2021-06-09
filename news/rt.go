package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RtNews() map[int]string {
	newsLinks := make(map[int]string)
	rt := "https://www.rt.com"
	res, err := http.Get(rt)

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	// Load the Html
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".news-block .main-promobox ul li .main-promobox__wrapper").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		newsLinks[i] = rt + href
		trim := strings.TrimSpace(title)
		fmt.Printf("Review %d: %s\n", i, trim)
	})

	return newsLinks
}
