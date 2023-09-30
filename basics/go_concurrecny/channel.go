package main

import (
	"fmt"
	"time"
)

func main() {
	//start := time.Now()

	myChannel := make(chan string, 1) //buffer(defined capacity) and un-buffer(0 capacity, needs a receiver as soon as something comes to channel.)

	myChannel <- "something"

	fmt.Println(<-myChannel)
	// go thisTakesForever("111")
	// go thisTakesForever("222")

	// fmt.Println("Done!")
	// fmt.Println(time.Since(start))

}

func thisTakesForever(id string) {
	//this does some work that takes a long time
	time.Sleep(time.Second * 2)
}
