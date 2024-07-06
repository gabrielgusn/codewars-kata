package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://softfocus.com.br",
		"http://mempool.space",
		"http://amazon.com",
		"http://hml-smart-infosimples.softfocus.com.br",
	}

	c := make(chan string)

	for _, url := range urls {
		syncCheckUrl(url)
	}

	for _, url := range urls {
		go asyncCheckUrl(url, c)
	}

	for url := range c {
		// time.Sleep(time.Second * 4)
		// go asyncCheckUrl(url, c)
		go func(copied_url string) {
			time.Sleep(time.Second * 4)
			asyncCheckUrl(copied_url, c)
		}(url)
		close(c)
	}
}

func asyncCheckUrl(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while accessing", url, "\n", err)
		c <- url
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println("The website", url, "is up!")
		c <- url
	} else {
		fmt.Println("The website", url, "is returning", resp.Status)
		c <- url
	}
}

func syncCheckUrl(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while accessing", url, "\n", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println("The website", url, "is up!")
	} else {
		fmt.Println("The website", url, "is returning", resp.Status)
	}
}
