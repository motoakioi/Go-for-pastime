package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/oov/audio/wave"
//	"github.com/mjibson/go-dsp/wav"
//	"github.com/mjibson/go-dsp/fft"
)

// Show WAV data format
func FmtDisplay(wfe *wave.WaveFormatExtensible){

	fmt.Println("...")
	fmt.Println(" Samplerate:", wfe.Format.SamplesPerSec)
	fmt.Println(" Channels  :", wfe.Format.Channels)
	fmt.Println(" Bits      :", wfe.Format.BitsPerSample)

}

// main
func main(){

	// File open and close
	file, erFile := os.Open(`./3octaves.wav`)
	defer file.Close()
	// In case of error
	if erFile != nil {
		errors.New("Can NOT open .wav file.")
	}

	// Read data from .wav file
	data, wfe, erData := wave.NewReader(file)
	// In case of error
	if erData != nil{
		errors.New("Can NOT read .wav data.")
	}
	// Show wave data format
	FmtDisplay(wfe)

	// Create buf for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wfe.Format.Channels); i++ {
		inTmp = append(inTmp, make([]float64, wfe.Format.SamplesPerSec))
	}

	// Read wave data from struct
	n, erN := data.ReadFloat64Interleaved(inTmp)
	if erN != nil{
		errors.New("Can Not read data form struct")
	}
	// This is wave data
	fmt.Println(inTmp[1])

}
