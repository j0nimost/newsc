package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Newser interface {
	GetNews() map[int] string
}
type NewsLoader struct {
	Url, Query string
}

func (n NewsLoader) GetNews() map[int] string {
	newsLinks := make(map[int]string)
	res, err := http.Get(n.Url)

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	// Load the Html
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	
	doc.Find(n.Query).Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		
		p := strings.HasPrefix(href, "https")

		if !p {
			newsLinks[i] = n.Url + href
		} else {
			newsLinks[i] = href
		}
		
		trim := strings.TrimSpace(title)
		fmt.Printf("Review %d: %s\n", i, trim)
	})

	return newsLinks
}