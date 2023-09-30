package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	myChannel := make(chan string, 6) //buffer(defined capacity) and un-buffer(0 capacity, needs a receiver as soon as something comes to channel.)

	go thisTakesForever("111", myChannel)
	go thisTakesForever("222", myChannel)
	go thisTakesForever("333", myChannel)
	go thisTakesForever("444", myChannel)
	go thisTakesForever("555", myChannel)
	go thisTakesForever("666", myChannel)

	for len(myChannel) < 6 { //this will make the main to wait for myChannel to have two messages on it.
		// which will make the main function blocked until both go routines finish.

	} // in go while loop is a for loop with some condition

	close(myChannel)

	fmt.Println("Done!")
	fmt.Println(time.Since(start))

}

func thisTakesForever(id string, myChannel chan<- string) {
	//this does some work that takes a long time
	time.Sleep(time.Second * 2)
	myChannel <- "Done" // this is the message go routines we defined will put on the channel
}
