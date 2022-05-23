package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				continue
			}
		}
	}()

	fmt.Println("Waiting forever... ")
	// while a channel not receive any value, it will wait
	<-forever
}
