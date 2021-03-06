package main

import (
	"flag"
	"fmt"
	"newscli/news"
	"time"
)

func main() {

	medialist := `
		 hn : Hacker News
		 cap : Capital Radio (KE)
		 aj : Aljazeera (International) 
		 rt : RT (International)
		 ctv : Citizen Tv(KE)
		 enca : eNCA (SA)
		 chtv : Channels Tv (NG)
		 nwr: Nairobi Wire (KE)
		`

	fmt.Println("\u001b[32m", `
		
  _        _______           _______         _______  _       _________
 ( (    /|(  ____ \|\     /|(  ____ \       (  ____ \( \      \__   __/
 |  \  ( || (    \/| )   ( || (    \/       | (    \/| (         ) (   
 |   \ | || (__    | | _ | || (_____  _____ | |      | |         | |   
 | (\ \) ||  __)   | |( )| |(_____  )(_____)| |      | |         | |   
 | | \   || (      | || || |      ) |       | |      | |         | |   
 | )  \  || (____/\| () () |/\____) |       | (____/\| (____/\___) (___
 |/    )_)(_______/(_______)\_______)       (_______/(_______/\_______/
																	   
 
		`, "\u001b[0m")

	fmt.Println(`
	~~ Get the news briefing from the following broadcasters ~~
		`, medialist)

	var (
		media    string
		newsLink map[int]string
		n        news.Newser
	)

	flag.StringVar(&media, "media", "",
		`Pass a media name eg`+medialist)
	flag.Parse()

	switch media {
	case "cap":
		n = news.NewsLoader{Url: "https://www.capitalfm.co.ke/news/", Query: ".zox-feat-right-wrap .zox-side-list-wrap section"}
		break
	case "aj":
		n = news.NewsLoader{Url: "https://www.aljazeera.com", Query: ".container .fte-featured__content-wrapper__right .fte-featured__right-inner-articles-wrapper .fte-featured__article-content"}
		break
	case "rt":
		n = news.NewsLoader{Url: "https://www.rt.com", Query: ".news-block .main-promobox ul li .main-promobox__wrapper"}
		break
	case "ctv":
		n = news.NewsLoader{Url: "https://citizentv.co.ke/", Query: ".main-story .more-election-stories div"}
		break
	case "enca":
		n = news.NewsLoader{Url: "https://www.enca.com", Query: ".view-latest-news .view-content .item-list ul li"}
		break
	case "chtv":
		n = news.NewsLoader{Url: "https://www.channelstv.com/", Query: ".news-fold .news_content_fold .news-list-item"}
		break
	case "nwr":
		n = news.NewsLoader{Url: "http://nairobiwire.com", Query: "#block-content .block2-small-holder .cat-block-post-title h3"}
		break
	case "hn":
		n = news.NewsLoader{Url: "https://news.ycombinator.com/", Query: "body #hnmain tbody tr table tbody .athing .title"}
	default:
		fmt.Println("\u001b[31m", "No Media House Specified or Not Found", "\u001b[0m")
		return
	}

	newsLink = n.GetNews() // pull the first load of news
	ticker := time.NewTicker(10 * time.Minute)

	for {

		select {

		case t := <-ticker.C:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\u001b[33m", "Time of refresh: ", t, "\u001b[0m")
			newsLink = n.GetNews()

		default:
			fmt.Println("Get a specific headline (pass the number): ")
			var reviewNo int
			if _, err := fmt.Scanf("%d\n", &reviewNo); err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			l, ok := newsLink[reviewNo]

			if !ok {
				fmt.Println("\u001b[31m", "News link not found :(", "\u001b[0m")
			}

			fmt.Printf("Review: %d: %s\n", reviewNo, l)

		}
	}

}
