package main

import (
	"fmt"
	"net/http"
	"time"
)

func webSiteChecker() {
	siteList := []string{
		"http://www.google.com",
		"http://www.bing.com",
		"http://www.yahoo.com",
		"http://www.facebook.com",
	}

	// Create a channel to store the results of the goroutines.
	c := make(chan string)

	for _, site := range siteList {
		// everytime you use the keyword go, a new goroutine is created
		go checkSite(site, c)
	}

	for message := range c {
		func(msg string) {
			time.Sleep(5 * time.Second)
			// calling another routine to print the message
			go checkSite(message, c)
		}(message)
	}
}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!")
		// communicate back to the main routine
		c <- site
	} else {
		fmt.Println(site, "is up!")
		// communicate back to the main routine
		c <- site
	}
}
