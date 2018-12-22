package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/oov/audio/wave"
//	"github.com/mjibson/go-dsp/wav"
//	"github.com/mjibson/go-dsp/fft"
)

func FmtDisplay(wfe *wave.WaveFormatExtensible){

	fmt.Println("...")
	fmt.Println(" Samplerate:", wfe.Format.SamplesPerSec)
	fmt.Println(" Channels  :", wfe.Format.Channels)
	fmt.Println(" Bits      :", wfe.Format.BitsPerSample)

}


func main(){

	// File open and close
	file, erFile := os.Open(`./3octaves.wav`)
	defer file.Close()
	// In case of error
	if erFile != nil {
		errors.New("Can NOT open .wav file.")
	}

	// Read WAV data
	data, wfe, erData := wave.NewReader(file)
	// In case of error
	if erData != nil{
		errors.New("Can NOT read .wav data.")
	}
	// Show WAV data format
	FmtDisplay(wfe)

	// Data handle
	inBuf := [][]float64{}
	data.ReadFloat64Interleaved(inBuf)
	//for i, val := range data.Data{
	//	fmt.Println(i, val)
	//}


}
