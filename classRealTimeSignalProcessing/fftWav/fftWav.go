package main

import (
	"errors"
	"fmt"
	"math"
	"math/cmplx"
	"os"

	"github.com/oov/audio/wave"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

// Show WAV data format
func FmtDisplay(wfe *wave.WaveFormatExtensible) {

	fmt.Println("...")
	fmt.Println(" Samplerate:", wfe.Format.SamplesPerSec)
	fmt.Println(" Channels  :", wfe.Format.Channels)
	fmt.Println(" Bits      :", wfe.Format.BitsPerSample)

}

// Get wav data from file
func getWavData(fileName string) ([][]float64, *wave.WaveFormatExtensible) {

	// File open and close
	file, erFile := os.Open(fileName)
	defer file.Close()
	// In case of error
	if erFile != nil {
		errors.New("Can NOT open .wav file.")
	}

	// Read data from .wav file
	data, wfe, erData := wave.NewReader(file)
	// In case of error
	if erData != nil {
		errors.New("Can NOT read .wav data.")
	}

	// Create buffer for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wfe.Format.Channels); i++ {
		inTmp = append(inTmp, make([]float64, wfe.Format.SamplesPerSec))
	}

	// Read wave data from struct
	data.ReadFloat64Interleaved(inTmp)

	return inTmp, wfe
}

// Create figure
func cre8Figure() *plot.Plot {

	fig, erFig := plot.New()
	if erFig != nil {
		errors.New("Can NOT create figure.")
	}

	return fig
}

// Set figure
func CfgFigure(fig *plot.Plot) {

	// Label config
	//fig.Title.Text = "CfgFigure func"
	fig.X.Label.Text = "x"
	fig.Y.Label.Text = "y"

	// Range for each axis
	fig.X.Min = 0
	fig.X.Max = 3000
	fig.Y.Min = -0.12
	fig.Y.Max = 0.12
}

// Set plot struct
func cfgPoint(x float64, dx float64, y []float64) plotter.XYs {
	plotTmp := make(plotter.XYs, int(x/dx))
	for i := 0; i < int(x/dx); i++ {
		plotTmp[i].X = float64(i) * dx
		plotTmp[i].Y = y[i]
	}
	return plotTmp
}

// Calculate power from complex number
func c2power(inC []complex128) []float64 {
	outR := []float64{}
	for i := 0; i < len(inC); i++ {
		outR = append(outR, math.Pow(cmplx.Abs(inC[1]), 2.0))
	}
	return outR
}

// main
func main() {

	// Get wav data from file
	wavData, wfe := getWavData("./3octaves.wav")

	// Show wave data format
	FmtDisplay(wfe)

	// !meaningless
	fmt.Println(wavData[0][0])
	/*
		// fft
		fftDataC := fft.FFTReal(wavData[0])

		// get power from complex number
		fftDataPow := c2power(fftDataC)
	*/
	// Create figure
	fig := cre8Figure()

	// Set figure
	CfgFigure(fig)

	//plotutil.AddLinePoints(fig, cfgPoint(float64(len(fftDataPow)), fftDataPow))
	plotutil.AddLinePoints(fig, "raw", cfgPoint(float64(len(wavData[0])), 1.0, wavData[0]))
	/*
		// Set function of plot
		plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
		plotFunc.Color = color.RGBA{B: 255, A: 255}
		fig.Add(plotFunc)
	*/
	// Save figure (width, height, file name)
	fig.Save(1500, 400, "wave.pdf")

	fmt.Println("Done.")

}

func myFunc(x float64) float64 {
	return 0.2*x + 50
}
