package main

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"math/cmplx"
	"os"

	"github.com/oov/audio/wave"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	// File Size
	fInfo, erFInfo := file.Stat()
	// In case of error
	if erFInfo != nil {
		errors.New("Can NOT get file information.")
	}
	fmt.Println("File size is ", fInfo.Size())

	// Read data from .wav file
	data, wfe, erData := wave.NewReader(file)
	// In case of error
	if erData != nil {
		errors.New("Can NOT read .wav data.")
	}

	// Calculate duration
	fSize := float32(fInfo.Size())
	wavCh := float32(wfe.Format.Channels)
	wavBit := float32(wfe.Format.BitsPerSample)
	numSample := int(fSize / (wavCh * (wavBit / 8.0)))
	fmt.Println("Duration is ", numSample)

	// Create buffer for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wavCh); i++ {
		inTmp = append(inTmp, make([]float64, numSample))
	}

	// Read wave data from struct
	n, erN := data.ReadFloat64Interleaved(inTmp)
	if erN != nil {
		errors.New("Can NOT read data from inTmp.")
	}
	fmt.Println("n : ", n)

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

type plotRange struct {
	xStart float64
	xEnd   float64
	yStart float64
	yEnd   float64
}

// Set figure
func CfgFigure(fig *plot.Plot, figRange plotRange) {

	// Label config
	//fig.Title.Text = "CfgFigure func"
	fig.X.Label.Text = "x"
	fig.Y.Label.Text = "y"

	// Range for each axis
	fig.X.Min = figRange.xStart
	fig.X.Max = figRange.xEnd
	fig.Y.Min = figRange.yStart
	fig.Y.Max = figRange.yEnd
}

// Set plot struct
func cfgPoint(x float64, dx float64, y []float64) plotter.XYs {
	plotTmp := make(plotter.XYs, int(x/dx))
	fmt.Println("x/dx : ", int(x/dx))
	for i := 0; i < int(x/dx); i++ {
		plotTmp[i].X = float64(i) * dx
		plotTmp[i].Y = y[i]
	}
	return plotTmp
}

// Add line from data
func addLine(fig *plot.Plot, xys plotter.XYs) {

	// Create Line
	plotLine, err := plotter.NewLine(xys)
	if err != nil {
		panic(err)
	}
	// Line config
	plotLine.LineStyle.Width = vg.Points(1)
	//plotLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	plotLine.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Add line to figure
	fig.Add(plotLine)
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
	wavData, wfe := getWavData("./1khz.wav")

	// Show wave data format
	FmtDisplay(wfe)

	/*
		// fft
		fftDataC := fft.FFTReal(wavData[0])

		// get power from complex number
		fftDataPow := c2power(fftDataC)
	*/

	// Create figure
	fig := cre8Figure()

	// Set range of plot
	var figRange plotRange
	figRange.xStart = 0
	figRange.xEnd = float64(len(wavData[0]))
	figRange.yStart = -1.5
	figRange.yEnd = 1.5

	// Set figure
	CfgFigure(fig, figRange)

	// Add data as line to figure
	fmt.Println("wav len : ", len(wavData[0]))
	addLine(fig, cfgPoint(float64(len(wavData[0])), 1.0, wavData[0]))

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
