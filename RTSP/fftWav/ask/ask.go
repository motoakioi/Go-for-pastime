package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/cmplx"
	"os"

	"github.com/oov/audio/wave"
//	"github.com/r9y9/go-dsp/fft"
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
	numSample := int(fSize / (wavCh * (wavBit / 8.0)))
	fmt.Println("Total sample points are ", numSample)

	// Create buffer for data handle
	inTmp := [][]float64{}
	for i := 0; i < int(wavCh); i++ {
		inTmp = append(inTmp, make([]float64, numSample))
	}

	// Read wave data from struct
	n, erN := data.ReadFloat64Interleaved(inTmp)
	if erN != nil {
		log.Fatal("Can NOT read data from inTmp.")
	}
	fmt.Println("n : ", n)

	return inTmp, wfe
}

// Create figure
func cre8Figure() *plot.Plot {

	fig, erFig := plot.New()
	if erFig != nil {
		log.Fatal("Can NOT create figure.")
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
	fig.X.Label.Text = "Sample"
	fig.Y.Label.Text = "Amplitude"

	// Range for each axis
	fig.X.Min = figRange.xStart
	fig.X.Max = figRange.xEnd
	fig.Y.Min = figRange.yStart
	fig.Y.Max = figRange.yEnd
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
		outR = append(outR, math.Pow(cmplx.Abs(inC[i]), 2.0))
	}
	return outR
}

func pow(in []float64) []float64{
	out := []float64{}
	for i:= 0; i<len(in); i++{
		out = append(out, math.Pow(in[i], 2.0))
	}
	return out
}

func datafunc(in []float64, size int) []float64{
	out := []float64{}
	for i:=0; i<(len(in)/size); i++{
		var tmp float64 = 0
		for j:=0; j<size; j++{
			tmp += float64(in[i*size+j])
		}
		if(tmp > 100){
			out = append(out, 1)
		}else{
			out = append(out, 0)
		}
	}
	return out
}

// main
func main() {

	// Get wav data from file
	wavData, wfe := getWavData("../wavData/ask.wav")

	// Show wave data format
	FmtDisplay(wfe)

	// Create Figure
	fig := cre8Figure()

	// Set range of plot
	var figRange plotRange
	figRange.xStart = 0
	figRange.xEnd = float64(len(wavData[0])/400)
	figRange.yStart = 0
	figRange.yEnd = 1

	// Set figure
	CfgFigure(fig, figRange)

	// Add data as line to figure
	addLine(fig, cfgPoint(figRange.xEnd, 1.0, datafunc(pow(wavData[0]), 400)))
	fmt.Println(datafunc(pow(wavData[0]), 400))

	/*
		// Set function of plot
		plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
		plotFunc.Color = color.RGBA{B: 255, A: 255}
		fig.Add(plotFunc)
	*/

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "data.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	fmt.Println("Done.")

}
