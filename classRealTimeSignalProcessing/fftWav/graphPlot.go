package main

import (
	"errors"
	"fmt"

	"gonum.org/v1/plot"
)

//main
func main() {

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

	// Save figure (width, height, file name)
	fig.Save(150, 150, "test.pdf")

	fmt.Println("Done")

}
