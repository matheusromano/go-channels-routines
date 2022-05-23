package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- "Hello World!"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("Default")
	}

	time.Sleep(time.Second * 3)
}
