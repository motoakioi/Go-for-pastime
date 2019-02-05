package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"strconv"

	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/figHandle"
	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/readWav"

	"github.com/r9y9/go-dsp/fft"
)

// Calculate power from complex number
func c2power(inC []complex128) []float64 {
	outR := []float64{}
	for i := 0; i < len(inC); i++ {
		valTmp := math.Pow(cmplx.Abs(inC[i]), 2.0)
		if valTmp < 0 {
			valTmp = -valTmp
		}
		outR = append(outR, math.Sqrt(valTmp))
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
	//fftData := make([]float64, size)
	for i := 0; i < times; i++ {
		tmpFftData[i] = myfft(wavData[0][(size * i):(size * (i + 1))])
		/*
			for j := 0; j < size; j++ {
				fftData[j] += tmpFftData[i][j]
			}
		*/
	}
	//fmt.Println(fftDataR)

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = 2000 // float64(size) / 8
	figRange.YStart = 0
	figRange.YEnd = 400

	// Set figure
	figHandle.CfgFigure(fig, figRange)

	// Add data as line to figure
	//figHandle.AddLine(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, fftData))
	tmp := 1
	for i := 5 * tmp; i < 5*(tmp+1)-1; i++ {
		var legend string
		legend += strconv.FormatFloat(float64(i)*0.5, 'g', 3, 64) + " - " + strconv.FormatFloat(float64(i+1)*0.5, 'g', 3, 64) + " s"
		figHandle.AddLineLegendColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, tmpFftData[i]), legend, (i+1)%5+i*5+1)
		fig.Legend.Top = true
	}

	// Save figure (width, height, file name)
	if fig.Save(300, 130, "piano2.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
