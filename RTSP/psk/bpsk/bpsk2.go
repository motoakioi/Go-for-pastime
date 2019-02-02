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
func DemBpsk(in []float64, period int) float64 {

	cosTb := []float64{}
	for i := 0; i < period; i++ {
		// Create cosine and sine tables to reference
		cosTb = append(cosTb, math.Cos(2*math.Pi*float64(i)/float64(period)))
	}
	d := 0.0      // sum of phase difference
	maxDif := 0.0 // value in case that phase difference is Pi
	for i := 0; i < len(in); i++ {
		d += math.Sqrt(math.Pow((cosTb[i%period] - in[i]), 2))
		maxDif += math.Sqrt(math.Pow((cosTb[i%period] * 2), 2))
	}
	fmt.Println(d / maxDif)
	return d / maxDif * math.Pi
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

func amplidata(in []float64) []float64 {
	out := []float64{}
	for i := 0; i < (len(in)); i++ {
		if in[i] <= 0 {
			out = append(out, -in[i])
		} else {
			out = append(out, in[i])
		}
	}
	return out
}

// main
func main() {

	// Get wav data from file
	wavData, wfe, _ := readWav.GetWavData("../wavData/bpsk2.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// Demodulation BPSK
	demData := []float64{}
	for i := 0; i < (len(wavData[0]) / 8); i++ {
		demData = append(demData, DemBpsk(wavData[0][i*8:i*8+8], 8))
		//demData = append(demData, DemBpsk(wavData[0][i*8:i*8+8], 8)...)
	}

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(len(demData))
	figRange.YStart = 0
	figRange.YEnd = 3.2

	// Set figure
	figHandle.CfgFigureName(fig, figRange, "Time [ms]", "Phase")

	// Add data as line to figure
	//figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0/8.0, amplidata(wavData[0])), 1)
	//figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0/8.0, DemBpsk(wavData[0], 8)), 10)
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, demData), 10)

	// Show domodulated data
	fmt.Println("Data : ", datafunc(demData, 400))

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "phase2.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
