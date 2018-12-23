package main

import (
	"errors"
	"fmt"
	"image/color"
	"os"

	"github.com/oov/audio/wave"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	//	"github.com/mjibson/go-dsp/wav"
	//	"github.com/mjibson/go-dsp/fft"
)

// Show WAV data format
func FmtDisplay(wfe *wave.WaveFormatExtensible) {

	fmt.Println("...")
	fmt.Println(" Samplerate:", wfe.Format.SamplesPerSec)
	fmt.Println(" Channels  :", wfe.Format.Channels)
	fmt.Println(" Bits      :", wfe.Format.BitsPerSample)

}

// Get wav data from file
func getWavData(fileName string) [][]float64 {

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

	// Show wave data format
	FmtDisplay(wfe)

	// Create buffer for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wfe.Format.Channels); i++ {
		inTmp = append(inTmp, make([]float64, wfe.Format.SamplesPerSec))
	}

	// Read wave data from struct
	data.ReadFloat64Interleaved(inTmp)

	return inTmp
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
	fig.X.Max = 200
	fig.Y.Min = 0
	fig.Y.Max = 200
}

// main
func main() {

	// Get wav data from file
	wavData := getWavData("./3octaves.wav")

	// !meaningless
	fmt.Println(wavData[0][0])

	// Create figure
	fig := cre8Figure()

	// Set figure
	CfgFigure(fig)

	// Set function of plot
	plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
	plotFunc.Color = color.RGBA{B: 255, A: 255}
	fig.Add(plotFunc)

	// Save figure (width, height, file name)
	fig.Save(150, 150, "test.pdf")

	fmt.Println("Done.")

}

func myFunc(x float64) float64 {
	return 0.2*x + 50
}
