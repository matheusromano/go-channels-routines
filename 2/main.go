package main

import "fmt"

// Thread 01
func main() {
	//Create a channel that will hold strings.
	// Thread 01 <-> Thread 02
	hello := make(chan string)

	// Thread 02
	//Start a goroutine that will send strings to the channel.
	go func() {
		//Send the strings.
		hello <- "Hello World!"
	}()

	//Everytime we receive a string from the channel, print it.
	result := <-hello
	fmt.Println(result)
}
