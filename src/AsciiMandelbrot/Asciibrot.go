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
	Color string
	ZReal float64
	ZComplex float64
}

// TODO: Create the map of color codes
var ColorMap = map[int]string
var MaxIters int = 0
var XScale float64 = 0.0
var YScale float64 = 0.0

func getColor(iters int, MaxIters int) string {
	iterVal := int(math.Round((float64(iters)/float64(MaxIters))*10))
	return ColorMap[iterVal]
}

func (this Pixel) getZValues(MaxIters int) {
	curIter := 0

	for ; math.Pow(this.ZReal, 2) + math.Pow(this.ZComplex, 2) <= 4.0) && (curIter < MaxIters); curIter++ {
		nextZReal := math.Pow(this.ZReal) - math.Pow(this.ZComplex) + float64(this.X)*YScale - 2
		nextZComplex := 2 * this.ZReal * this.ZComplex + float64(this.Y)*YScale - 2

		this.ZReal = nextZReal
		this.ZComplex = nextZComplex
	}

	this.Color = getColor(curIter, MaxIters)
}

// TODO: Make this a debug function, 
// Create a string function to only create a new string with a background color code, a space character, then a console-color default code.
func (this Pixel) String() string {
	return fmt.Sprintf("(%d, %d), (%f, %f), Color:%s ", this.X, this.Y, this.ZReal, this.ZComplex, this.Color)
}

// TODO: Add a thread-block before starting to print all the pixel info.
// TODO: add the multithreading back in, & limit logical threads to number of CPU threads.
// TODO: Fix for each Pixel printing.
// TODO: Make the cmdline arg parsing better.
// TODO: more todos
func main() {
	fmt.Println("")
	cmdArgs := os.Args[1:]

	xSize, ex := strconv.Atoi(cmdArgs[0])
	ySize, ey := strconv.Atoi(cmdArgs[1])
	iters, ei := strconv.Atoi(cmdArgs[2])
	
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
	
	pixels := make([]Pixel, x*y)
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
			index := x + y%x;
			pixels[index] = Pixel{x, y, "", 0.0, 0.0}
			pixels[index].getZValues(iters)
		}
	}

	for pix := range pixels {
		fmt.Println(pix)
	}
}