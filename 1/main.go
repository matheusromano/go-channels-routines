package main

import (
	"fmt"
	"time"
)

func count(dataType string) {
	for i := 0; i < 5; i++ {
		fmt.Println(dataType, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 	count("Without go routine")
	// 	// this line of code will execute in a concurrent manner
	// 	go count("With go routine")
	// 	fmt.Println("Hello Earth")
	// 	fmt.Println("Hello Jupyter")
	// 	fmt.Println("..finished")
	//

	// this line of code will execute in a concurrent manner
	go count("first")
	go count("second")
	time.Sleep(5 * time.Second)
}
