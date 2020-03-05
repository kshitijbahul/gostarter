package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	c := make(chan string) //creating a new Channel

	for _, link := range links {
		go checkRequest(link, c) //the main routine will end not waiting for the child routines to finish

	}
	//for i := 0; i < len(links); i++ {
	//for { //value into the channel
	for l := range c { //ruynning through the channel everytime the channel emmits a value
		//fmt.Println(<-c)
		//go checkRequest(<-c, c)

		//go checkRequest(l, c) //commented to add a limit of 5 seconds
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkRequest(link, c)
		}(l) //extra () are to call the function

	} //This might be about recieveing the message from the channels

}

func checkRequest(url string, c chan string) {
	_, error := http.Get(url)
	if error != nil {
		fmt.Println(url, " might be down")
		//c <- "Might be Down "
		c <- url
		return
	}
	fmt.Println(url, " is up")
	//c <- "Its up "
	c <- url
}
