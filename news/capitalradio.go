package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CapitalRadio struct {
	Url, Query string
}

func (cr CapitalRadio) GetNews() map[int] string {
	news := make(map[int]string)

	res, err := http.Get(cr.Url)
	if err != nil {
		log.Fatalf("Status Code Error: %d : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(cr.Query).Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		href := s.Find("a").AttrOr("href", "")
		news[i] = href
		trimed := strings.TrimSpace(title)
		fmt.Printf("Review %d: %s\n", i, trimed)
	})

	return news
}