package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Started ...")

	go func() {
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Finished ...")

}
