package figHandle

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Create figure
func Cre8Figure() *plot.Plot {

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
func CfgPoint(x float64, dx float64, y []float64) plotter.XYs {
	plotTmp := make(plotter.XYs, int(x/dx))
	for i := 0; i < int(x/dx); i++ {
		plotTmp[i].X = float64(i) * dx
		plotTmp[i].Y = y[i]
	}
	return plotTmp
}

// Add line from data
func AddLine(fig *plot.Plot, xys plotter.XYs) {

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
