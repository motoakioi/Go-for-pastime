package readWav

import (
	"fmt"
	"log"
	"os"

	"github.com/oov/audio/wave"
)

// Show WAV data format
func FmtDisplay(wfe *wave.WaveFormatExtensible) {

	fmt.Println("...")
	fmt.Println(" Samplerate:", wfe.Format.SamplesPerSec)
	fmt.Println(" Channels  :", wfe.Format.Channels)
	fmt.Println(" Bits      :", wfe.Format.BitsPerSample)

}

// Get wav data from file
func GetWavData(fileName string) ([][]float64, *wave.WaveFormatExtensible, int32) {

	// File open and close
	file, erFile := os.Open(fileName)
	defer file.Close()
	// In case of error
	if erFile != nil {
		log.Fatal("Can NOT open .wav file.")
	}

	// File Size
	fInfo, erFInfo := file.Stat()
	// In case of error
	if erFInfo != nil {
		log.Fatal("Can NOT get file information.")
	}
	fmt.Println("File size is ", fInfo.Size())

	// Read data from .wav file
	data, wfe, erData := wave.NewReader(file)
	// In case of error
	if erData != nil {
		log.Fatal("Can NOT read .wav data.")
	}

	// Calculate duration
	fHeaderSize := int64(44)
	fSize := float32(fInfo.Size() - fHeaderSize)
	wavCh := float32(wfe.Format.Channels)
	wavBit := float32(wfe.Format.BitsPerSample)
	numSample := int32(fSize / (wavCh * (wavBit / 8.0)))
	fmt.Println("Total sample points are ", numSample)

	// Create buffer for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wavCh); i++ {
		inTmp = append(inTmp, make([]float64, numSample))
	}

	// Read wave data from struct
	_, erN := data.ReadFloat64Interleaved(inTmp)
	if erN != nil {
		log.Fatal("Can NOT read data from inTmp.")
	}

	return inTmp, wfe, numSample
}
