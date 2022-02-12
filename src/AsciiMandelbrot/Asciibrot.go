package main

import (
	"os"
	"fmt"
	"math"
	"strconv"
)

type Pixel struct {
	X int
	Y int
	Color *string
	ZReal float64
	ZComplex float64
}

// TODO: Create the map of color codes
var ColorMap = []string{"40","41","42","43","44","45","46","47","48","49","50"}

var MaxIters int = 0
var XScale float64 = 0.0
var YScale float64 = 0.0

func (this *Pixel) getColor(iters int) {
	this.Color = &ColorMap[ int( math.Floor((float64(iters) / float64(MaxIters)) * 10) )]
}

func qsquare(base float64) float64 {
	return base * base
}

func (this *Pixel) getZValues() {
	curIter := 0
	
	zReal0 := float64(this.X)*XScale - 2
	zComplex0 := float64(this.Y) * YScale - 2

	for ; (qsquare(this.ZReal) + qsquare(this.ZComplex) <= 4.0) && (curIter < MaxIters); curIter++ {
		nextZReal := qsquare(this.ZReal) - qsquare(this.ZComplex) + zReal0
		nextZComplex := 2 * this.ZReal * this.ZComplex + zComplex0

		this.ZReal = nextZReal
		this.ZComplex = nextZComplex
	}

	this.getColor(curIter)
}

// TODO:
func (this Pixel) String() string {
	return fmt.Sprintf("\033[%sm \033[0m", *this.Color)
}

// TODO: Add a thread-block before starting to print all the pixel info.
// TODO: add the multithreading back in, & limit logical threads to number of CPU threads.
// TODO: Make the cmdline arg parsing better.
func main() {
	fmt.Println("")
	cmdArgs := os.Args[1:]

	xSize, ex := strconv.Atoi(cmdArgs[0])
	ySize, ey := strconv.Atoi(cmdArgs[1])
	var ei error = nil; // Fixes scoping problem for MaxIters
	MaxIters, ei = strconv.Atoi(cmdArgs[2])
	
	if (ex != nil || ey != nil || ei != nil) {
		if (ex != nil) {
			fmt.Println("Incompatible x argument")
		}
		
		if (ey != nil) {
			fmt.Println("Incompatible y argument")
		}

		if (ei != nil) {
			fmt.Println("Incompatible iters argument")
		}

		fmt.Println("Exiting...")
		return
	}

	pixels := make([]Pixel, xSize*ySize)
	// x/xtot = mx/mtot, where mtot = 4 (range of mandelbrot set)
	// -> x * mtot/xtot = mx
	// This gives us the scale, but not the exact value of zreal0.
	// Because the (min, max) of the mandelbrot == (-2.0, 2.0) for each variable,
	// we need to subtract 2 to get to the "true" scaled value.
	// Thus, we get:
	// zreal_0 = (x * mtot/xtot) - 2
	// This follows for y and zcomplex_0
	XScale = 4.0/float64(xSize)
	YScale = 4.0/float64(ySize)

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			index := x + y*xSize;
			pixels[index] = Pixel{x, y, &ColorMap[0], 0.0, 0.0}
			pixels[index].getZValues()
		}
	}

	for _, pix := range pixels {
		fmt.Printf("%s", pix)
		if pix.X == xSize - 1 {
			fmt.Print("\n")
		}
	}
}