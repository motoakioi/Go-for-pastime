package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"

	"./figHandle"
	"./readWav"

	"github.com/r9y9/go-dsp/fft"
)

/*
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
	_, erN := data.ReadFloat64Interleaved(inTmp)
	if erN != nil {
		log.Fatal("Can NOT read data from inTmp.")
	}

	return inTmp, wfe
}
*/

/*
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
	fig.X.Label.Text = "Frequency [Hz]"
	//fig.Y.Label.Text = "y"

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
*/

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
	wavData, wfe := readWav.GetWavData("../wavData/3octaves.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// fft
	fftDataC := []complex128{}
	fftDataC = fft.FFTReal(wavData[0])

	// get power from complex number
	fftDataPow := c2power(fftDataC)

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange plotRange
	figRange.xStart = 0
	//figRange.xEnd = float64(len(fftDataPow) / 2)
	figRange.xEnd = float64(wfe.Format.SamplesPerSec / 2)
	figRange.yStart = 0
	figRange.yEnd = 1500000

	// Set figure
	figHandle.CfgFigure(fig, figRange)

	fmt.Println("x range after", figRange.xEnd)
	// Add data as line to figure
	figHandle.addLine(fig, cfgPoint(figRange.xEnd, 1.0, fftDataPow))

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
