package main

import (
	"fmt"
	"log"
)

const (
	waitTime = 60
)

func aMinuteWait(outCh chan string, num int) {
	//time.Sleep(waitTime * 1000 * 1000 * time.Microsecond)
	var a int
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 1000000000; j++ {
			//if a < i {
			a = i
			//}
		}
	}

	//log.Print("Waited a minute.")
	fmt.Println(num, "Waited a minute.")
	outCh <- "finished waiting for a minute\n"
}

func main() {
	outCh := make(chan string)
	log.Print("Starting...")
	go aMinuteWait(outCh, 1)
	go aMinuteWait(outCh, 2)
	go aMinuteWait(outCh, 3)
	go aMinuteWait(outCh, 4)
	go aMinuteWait(outCh, 5)
	go aMinuteWait(outCh, 6)
	msg1, msg2, msg3, msg4, msg5, msg6 := <-outCh, <-outCh, <-outCh, <-outCh, <-outCh, <-outCh
	log.Print("Done.")
	fmt.Println(msg1, msg2, msg3, msg4, msg5, msg6)
}
