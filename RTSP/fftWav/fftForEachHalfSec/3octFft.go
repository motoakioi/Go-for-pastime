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

// main
func main() {

	// Get wav data from file
	wavData, wfe, totalSample := readWav.GetWavData("../wavData/3octaves.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// fft
	duration := 0.5
	size := int(wfe.Format.SamplesPerSec * duration)
	times := totalSample / size
	fftDataR := make([]float64, size)
	for i := 0; i < times; i++ {
		tmpFftC := fft.FFTReal(wavData[0][(size * i):(size*(i+1) - 1)])
		fftDataR += c2power(tmpFftC)
	}
	//fftDataC := fft.FFTReal(wavData[0])

	// get power from complex number
	//fftDataPow := c2power(fftDataC)

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(wfe.Format.SamplesPerSec / 2)
	figRange.YStart = 0
	figRange.YEnd = 1500000

	// Set figure
	figHandle.CfgFigure(fig, figRange)

	fmt.Println("x range after", figRange.XEnd)
	// Add data as line to figure
	//figHandle.AddLine(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, fftDataPow))
	figHandle.AddLine(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, fftDataR))

	/*
		// Set function of plot
		plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
		plotFunc.Color = color.RGBA{B: 255, A: 255}
		fig.Add(plotFunc)
	*/

	// Save figure (width, height, file name)
	if fig.Save(1000, 400, "fft.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
