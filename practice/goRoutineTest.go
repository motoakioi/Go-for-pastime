package main

import (
	"fmt"
	"log"
	"time"
)

const (
	waitTime = 60
)

func aMinuteWait() {
	time.Sleep(waitTime * 1000 * 1000 * time.Microsecond)
	//log.Print("Waited a minute.")
	fmt.Println("Waited a minute.")
}

func main() {
	log.Print("Starting...")
	go aMinuteWait()
	go aMinuteWait()
	aMinuteWait()
	log.Print("Done.")
}
