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

	flag.StringVar(&media, "media", "", "Pass a media name eg cap")
	flag.Parse()

	if media == "cap" {
		news.Capitalradio()
	}

}
