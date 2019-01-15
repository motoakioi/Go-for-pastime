package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"

	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/figHandle"
	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/readWav"

	"github.com/r9y9/go-dsp/fft"
)

// Calculate power from complex number
func c2power(inC []complex128) []float64 {
	outR := []float64{}
	for i := 0; i < len(inC); i++ {
		outR = append(outR, math.Pow(cmplx.Abs(inC[i]), 2.0))
	}
	return outR
}

func myfft(in []float64) []float64 {
	tmpFftC := fft.FFTReal(in)
	return c2power(tmpFftC)
}

// main
func main() {

	// Get wav data from file
	wavData, wfe, totalSample := readWav.GetWavData("../wavData/3octaves.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// fft
	duration := float32(0.5)
	size := int(float32(wfe.Format.SamplesPerSec) * duration)
	fmt.Println("size", size)
	times := int(float32(totalSample) / float32(size))
	tmpFftData := make([][]float64, times, size)
	fftData := make([]float64, size)
	for i := 0; i < times; i++ {
		tmpFftData[i] = myfft(wavData[0][(size * i):(size * (i + 1))])
		for j := 0; j < size; j++ {
			fftData[j] += tmpFftData[i][j]
		}
	}
	//fmt.Println(fftDataR)

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(size) / 8
	figRange.YStart = 0
	figRange.YEnd = 150000

	// Set figure
	figHandle.CfgFigure(fig, figRange)

	fmt.Println("x range ", figRange.XEnd)
	fmt.Println("data len ", len(fftData))
	// Add data as line to figure
	figHandle.AddLine(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, fftData))

	/*
		// Set function of plot
		plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
		plotFunc.Color = color.RGBA{B: 255, A: 255}
		fig.Add(plotFunc)
	*/

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "piano.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
