package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UrlInfo struct {
	Url  string
	Size int
}

func main() {
	daChannel := make(chan UrlInfo)
	urls := [6]string{
		"https://example.com",
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"http://www.codeminders.com/weather_similarity/",
		"https://espn.com",
	}
	for _, url := range urls {
		go responseSize(url, daChannel)
	}
	for i := 0; i < len(urls); i++ {
		pageInfo := <-daChannel
		fmt.Println("url", pageInfo.Url, "is", pageInfo.Size, "bytes big")
	}
}

func responseSize(url string, channel chan UrlInfo) {
	fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- UrlInfo{
		Url:  url,
		Size: len(body),
	}
}
