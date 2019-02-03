package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/figHandle"
	"github.com/motoakioi/Go-for-pastime/RTSP/fftWav/fftForEachHalfSec/readWav"
)

// BPSK demodulation
// "in" is input signal, "period" is number of sample per cycle,
// "maxAmp" is maximun amplitude of input signal,
// if "initialPhase" =0.0, create cosine table, if =pi/2, create sine table
func DemBpsk(in []float64, period int, maxAmp float64, initialPhase float64) float64 {

	Tb := []float64{}
	for i := 0; i < period; i++ {
		// Create Cosine or Sine tables to reference,
		// Used Sine, but able to create Cosine table by changing "initialPhase"
		Tb = append(Tb, math.Sin(2*math.Pi*float64(i)/float64(period)+initialPhase))
	}
	d := 0.0      // sum of phase difference
	maxDif := 0.0 // value in case that phase difference is Pi

	for i := 0; i < len(in); i++ {
		d += math.Sqrt(math.Pow((maxAmp*Tb[i%period] - in[i]), 2))
		maxDif += math.Sqrt(math.Pow((maxAmp * Tb[i%period] * 2), 2))
	}
	return d / maxDif * math.Pi
}

// QPSK demodulation
// This function uses BPSK function twice
// "in" is input signal, "period" is number of sample per cycle,
// "channel" is number of channel of input data,
// "position" is demodulation position
func DemQpsk(in [][]float64, period int, channel int, position int) []float64 {

	demData := []float64{0, 0}
	for ch := 0; ch < channel; ch++ {
		CosineOrSine := float64(math.Pi * float64(ch%channel) / 2.0)
		demData[ch] = DemBpsk(in[ch][position*8:position*8+8], 8, maxValue(in[ch]), CosineOrSine)
	}
	return demData
}

// Return max value of slice
func maxValue(in []float64) float64 {
	max := 0.0
	for i := 0; i < len(in); i++ {
		if max < in[i] {
			max = in[i]
		}
	}
	// In case of max value of input slice is 0,
	// can not calculate power in demodulate function.
	if max == 0.0 {
		max = 0.01
	}
	return max
}

// Judge binary from phase
func datafunc(in []float64, size int) []float64 {
	out := []float64{}
	maxValTmp := maxValue(in) * float64(size) / 2
	for i := 0; i < (len(in) / size); i++ {
		tmp := 0.0
		for j := 0; j < size; j++ {
			tmp += float64(in[i*size+j])
		}
		if tmp > maxValTmp {
			out = append(out, 0)
		} else {
			out = append(out, 1)
		}
	}
	return out
}

// Show message encoding from binary ASCII data
func data2msg(in [2][]float64) {

	var msgASCII string
	for i := 1; i < len(in[0]); i++ {

		for j := 0; j < len(in); j++ {
			// Add binary to tmp
			msgASCII += strconv.FormatFloat(in[j][i], 'G', 1, 64)
		}

		// Insert space for each 8 bit (8 = 4 x channel)
		if i%4 == 0 {
			msgASCII += " "
		}
	}
	fmt.Println(msgASCII)

	var msg string
	for i := 0; i < len(msgASCII)/9; i++ {

		// Convert to decimal
		decimal, _ := strconv.ParseUint(msgASCII[i*9:i*9+8], 2, 10)

		// Convert to letter
		msg += string(decimal)
	}
	// Show whole message
	fmt.Println("Message is '", msg, "'")
}

// main
func main() {

	// Get wav data from file
	wavData, wfe, _ := readWav.GetWavData("../wavData/qpsk.wav")

	// Show wave data format
	readWav.FmtDisplay(wfe)

	// Demodulation QPSK
	windowSize := 8
	channel := int(wfe.Format.Channels)
	demData := [2][]float64{}
	for i := 0; i < (len(wavData[0]) / windowSize); i++ {

		// Demodulation for each windowsize (at position i to i+7)
		demDataTmp := DemQpsk(wavData, windowSize, channel, i)

		// Split data to each channel
		for ch := 0; ch < channel; ch++ {
			demData[ch] = append(demData[ch], demDataTmp[ch])
		}

	}

	// Create figure
	fig := figHandle.Cre8Figure()

	// Set range of plot
	var figRange figHandle.PlotRange
	figRange.XStart = 0
	figRange.XEnd = float64(len(demData[0]))
	figRange.YStart = 0
	figRange.YEnd = 3.2

	// Set figure
	figHandle.CfgFigureName(fig, figRange, "Time [ms]", "Phase")

	// Add data as line to figure
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, demData[0]), 10)
	figHandle.AddLineColor(fig, figHandle.CfgPoint(figRange.XEnd, 1.0, demData[1]), 1)

	// Save figure (width, height, file name)
	if fig.Save(500, 200, "test0.pdf") != nil {
		log.Fatal("Can NOT save figure.")
	}

	// Show domodulated data
	binData := [2][]float64{}
	binData[0] = append(binData[0], datafunc(demData[0], 50)...)
	binData[1] = append(binData[1], datafunc(demData[1], 50)...)
	fmt.Println(" Left : ", binData[0])
	fmt.Println(" Right : ", binData[1])

	// convert binary data to message
	data2msg(binData)

	fmt.Println("Done.")

}
