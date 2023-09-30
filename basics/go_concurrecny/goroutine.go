package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go thisTakesForever("111") // we can't move on to next until this is done.(using go keyword we can actually avoid waiting)
	go thisTakesForever("222") // we can't move on to next until this is done.(using go keyword we can actually avoid waiting)

	fmt.Println("Done!")
	fmt.Println(time.Since(start))

}

func thisTakesForever(id string) {
	//this does some work that takes a long time
	time.Sleep(time.Second * 2)
}
