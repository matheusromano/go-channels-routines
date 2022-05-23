package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("Executing Goroutine")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("Goroutine finished")
}
func checkChannels() {
	fmt.Println("Go channels are awesome!")

	values := make(chan string, 2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	value := <-values
	fmt.Println(value)

	time.Sleep(1 * time.Second)

}
