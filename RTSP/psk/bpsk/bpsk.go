package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"

	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/figHandle"
	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/readWav"
)

// Calculate power from complex number
func c2power(inC []complex128) []float64 {
	outR := []float64{}
	for i := 0; i < len(inC); i++ {
		outR = append(outR, math.Pow(cmplx.Abs(inC[i]), 2.0))
	}
	return outR
}
func pow(in []float64) []float64 {
	out := []float64{}
	for i := 0; i < len(in); i++ {
		out = append(out, math.Pow(in[i], 2.0))
	}
	return out
}
func add(in1 []float64, in2 []float64) []float64 {
	out := []float64{}
	for i := 0; i < len(in1); i++ {
		out = append(out, in1[i]+in2[i])
	}
	return out
}

// BPSK demodulation
// "in" is input signal, "period" is number of sample per cycle
func DemBpsk(in []float64, period int) []float64 {
	cosTb := []float64{}
	sinTb := []float64{}
	for i := 0; i < period; i++ {
		// Create cosine and sine tables to reference
		cosTb = append(cosTb, math.Cos(2*math.Pi*float64(i)/float64(period)))
		sinTb = append(sinTb, math.Sin(2*math.Pi*float64(i)/float64(period)))
	}
	ixc := []float64{}
	ixs := []float64{}
	for i := 0; i < len(in); i++ {
		ixc = append(ixc, cosTb[i%period]*in[i])
		ixs = append(ixs, sinTb[i%period]*in[i])
	}
	return add(ixc, ixs)
}

func datafunc(in []float64, size int) []float64 {
	out := []float64{}
	for i := 0; i < (len(in) / size); i++ {
		var tmp float64 = 0
		for j := 0; j < size; j++ {
			tmp += float64(in[i*size+j])
		}
		if tmp > 100 {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}

// main
func main() {

	// Get wav data from file
	wavData, wfe, _ := readWav.GetWavData("../wavData/bpsk1.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// Demodulation BPSK
	demData := DemBpsk(wavData[0], 8)

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(len(wavData[0]))
	figRange.YStart = -1.1
	figRange.YEnd = 1.1

	// Set figure
	figHandle.CfgFigureName(fig, figRange, "Sample", "Sum of products")

	// Add data as line to figure
	//figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, wavData[0]), 1)
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, DemBpsk(wavData[0], 8)), 10)

	// Show domodulated data
	fmt.Println("Data : ", datafunc(demData, 400))

	/*
		// Set function of plot
		plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
		plotFunc.Color = color.RGBA{B: 255, A: 255}
		fig.Add(plotFunc)
	*/

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "wave.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
