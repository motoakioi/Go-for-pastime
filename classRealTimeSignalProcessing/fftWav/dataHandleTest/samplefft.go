package main

import (
	"fmt"
	"log"
	"os"
	"github.com/r9y9/go-dsp/wav"
)

func main() {
	// ファイルのオープン
	file, err := os.Open("./3octaves.wav")
	if err != nil {
		log.Fatal(err)
	}

	// Wavファイルの読み込み 
	w, werr := wav.ReadWav(file)
	if werr != nil {
		log.Fatal(werr)
	}

	// データを表示
	for i, val := range w.Data {
		fmt.Println(i, val)
	}
}
