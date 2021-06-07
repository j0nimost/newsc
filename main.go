package main

import (
	"flag"
	"fmt"
	"newscli/news"
)

func main() {

	fmt.Println(`
		
  _        _______           _______         _______  _       _________
 ( (    /|(  ____ \|\     /|(  ____ \       (  ____ \( \      \__   __/
 |  \  ( || (    \/| )   ( || (    \/       | (    \/| (         ) (   
 |   \ | || (__    | | _ | || (_____  _____ | |      | |         | |   
 | (\ \) ||  __)   | |( )| |(_____  )(_____)| |      | |         | |   
 | | \   || (      | || || |      ) |       | |      | |         | |   
 | )  \  || (____/\| () () |/\____) |       | (____/\| (____/\___) (___
 |/    )_)(_______/(_______)\_______)       (_______/(_______/\_______/
																	   
 
		`)

	fmt.Println(`
		~~ Get the news briefing from the following broadcasters ~~\n
		1. cap : Capital Radio (KE)
		`)

	var media string
	var newsLink map[int]string

	flag.StringVar(&media, "media", "", "Pass a media name eg cap")
	flag.Parse()

	if media == "cap" {
		newsLink = news.Capitalradio()
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
