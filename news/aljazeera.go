package news

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Aljazeera struct {
	Url, Query string
}

func (aj Aljazeera) GetNews() map[int] string {
	news := make(map[int]string)
	res, err := http.Get(aj.Url)

	if err != nil {
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	
	doc.Find(aj.Query).Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		news[i] = aj.Url + href
		fmt.Printf("Review %d: %s\n", i, title)
	})

	return news
}