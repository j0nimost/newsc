package main

import (
	"flag"
	"fmt"
	"newscli/news"
)

func main() {

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
		1. cap : Capital Radio (KE)
		2. aj : Aljazeera (International) 
		3. rt : RT (International)
		4. ctv : Citizen Tv(KE)
		`)

	var (
		media string
		newsLink map[int]string
		n news.Newser
	)

	flag.StringVar(&media, "media", "",
		`Pass a media name eg
			1. cap : Capital Radio (KE)
			2. aj : Aljazeera (International) 
			3. rt : RT (International)	
			4. ctv : Citizen Tv(KE)`)
	flag.Parse()

	if media == "cap" {
		n = news.NewsLoader{Url: "https://www.capitalfm.co.ke/news/", Query:".zox-feat-right-wrap .zox-side-list-wrap section"}
		newsLink = n.GetNews()
	} else if media == "aj" {
		n = news.NewsLoader{Url: "https://www.aljazeera.com", Query: ".container .fte-featured__content-wrapper__right .fte-featured__right-inner-articles-wrapper .fte-featured__article-content"}
		newsLink = n.GetNews()
	} else if media == "rt" {
		n = news.NewsLoader{Url: "https://www.rt.com", Query:".news-block .main-promobox ul li .main-promobox__wrapper"}
		newsLink = n.GetNews()
	} else if media == "ctv" {
		n = news.NewsLoader{Url: "https://citizentv.co.ke/", Query:".main-story .more-election-stories div"}
		newsLink = n.GetNews()
	} else {
		fmt.Println("No Media House Specified")
		return
	}

	for {
		fmt.Println("Get a specific headline (pass the number): ")
		var reviewNo int
		if _, err := fmt.Scanf("%d\n", &reviewNo); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		l, ok := newsLink[reviewNo]

		if !ok {
			fmt.Println("News link not found :(")
			return
		}

		fmt.Printf("Review: %d: %s\n", reviewNo, l)

	}

}
