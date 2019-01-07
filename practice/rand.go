package main 

import (
	"fmt"
	"errors"
	"os"
	"math/rand"
	"time"
	"strconv"
)

func main(){

	if len(os.Args)!=2{
		fmt.Println("Usage : ", os.Args[0], "[number]")
	}
	inNumber, erInNumber := strconv.Atoi(os.Args[1])
	fmt.Println("input number", os.Args[1])

	if erInNumber != nil {
		errors.New("Can NOT convert from Args to int.")
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random Number is ", rand.Intn(inNumber))
}
