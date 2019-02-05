package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"

	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/figHandle"
	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/readWav"
	"github.com/oov/audio/wave"

	"github.com/r9y9/go-dsp/fft"
)

// Calculate power from complex number
func c2power2sqrt(inC []complex128) []float64 {
	outR := []float64{}
	for i := 0; i < len(inC); i++ {
		outR = append(outR, math.Sqrt(math.Pow(cmplx.Abs(inC[i]), 2.0)))
	}
	return outR
}

func myfft(in []float64) []float64 {
	tmpFftC := fft.FFTReal(in)
	return c2power2sqrt(tmpFftC)
}

// main
func main() {

	// Get wav data from file
	wavData, wfe, totalSample := readWav.GetWavData("./my8.wav")

	// Put data to file
	file, erFile := os.Create("out.wav")
	// In case of error
	if erFile != nil {
		log.Fatal("Can NOT create .wav file.")
	}
	defer file.Close()

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// up sampling
	magnification := 4 // 8 kHz to 32 KHz
	upsampled := []float64{}
	for _, d := range wavData[0] {
		for j := 0; j < magnification; j++ {
			if j == 0 {
				upsampled = append(upsampled, d)
			} else {
				upsampled = append(upsampled, 0)
			}
		}
	}

	newWfe := wave.WaveFormatExtensible{wfe.Format, wfe.Samples, wfe.ChannelMask, wfe.SubFormat}
	newWfe.Format.SamplesPerSec = wfe.Format.SamplesPerSec * uint32(magnification)
	outData := make([][]float64, newWfe.Format.Channels)
	fmt.Println("ori : ", wfe.Format.SamplesPerSec, ", new : ", newWfe.Format.SamplesPerSec, ", cal : ", wfe.Format.SamplesPerSec*uint32(magnification))
	for i := 0; i < int(newWfe.Format.Channels); i++ {
		outData[i] = make([]float64, len(upsampled))
	}
	for i := range upsampled {
		outData[0][i] = upsampled[i]
	}

	//fft
	fftData := myfft(upsampled)[0 : len(wavData[0])*int(magnification)/2]

	// LPF
	cutOffFreq := 4000 // Hz
	lpfData := []float64{}
	for i, d := range fftData {
		if i < cutOffFreq {
			lpfData = append(lpfData, d)
		} else {
			lpfData = append(lpfData, 0)
		}
	}

	//aw, err := wave.NewWriter(file, newWfe)
	aw, err := wave.NewWriter(file, &newWfe)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aw.Close()

	_, err = aw.WriteFloat64Interleaved(outData)
	if err != nil {
		fmt.Println("error")
		return
	}

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(len(fftData))
	figRange.YStart = 0
	figRange.YEnd = 150

	// Set figure
	figHandle.CfgFigureName(fig, figRange, "Frequency [Hz]", "Amplitude")

	// Add data as line to figure
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, fftData), 1)
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, lpfData), 10)

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "up-lpf.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
