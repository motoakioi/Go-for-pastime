package main

import (
	"errors"
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

//main
func main() {

	// Create figure
	fig, erFig := plot.New()
	if erFig != nil {
		errors.New("Can NOT create figure")
	}

	// Label config
	//fig.Title.Text = "test Fig"
	fig.X.Label.Text = "x"
	fig.Y.Label.Text = "y"

	// Range for each axis
	fig.X.Min = 0
	fig.X.Max = 100
	fig.Y.Min = 0
	fig.Y.Max = 100

	// Set function of plot
	plotFunc := plotter.NewFunction(func(x float64) float64 { return myFunc(x) })
	plotFunc.Color = color.RGBA{B: 255, A: 255}

	fig.Add(plotFunc)

	// Save figure (width, height, file name)
	fig.Save(150, 150, "test.pdf")

	fmt.Println("Done")

}

func myFunc(x float64) float64 {
	return x
}
