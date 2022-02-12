package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

type Pixel struct {
	X        int
	Y        int
	Color    *string
	ZReal    float64
	ZComplex float64
}

type ViewPort struct {
	XMin float64
	XMax float64
	YMin float64
	YMax float64
}

var waitgroup sync.WaitGroup

// TODO: Create the map of color codes
var ColorMap = []string{"40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50"}

var MaxIters int = 0
var XScale float64 = 0.0
var YScale float64 = 0.0
var FractalArea = ViewPort{-2.0, 2.0, -2.0, 2.0}

// TODO: Improve this for larger iteration counts.
func (this *Pixel) getColor(iters int) {
	this.Color = &ColorMap[int(math.Floor((float64(iters)/float64(MaxIters))*10))]
}

func qsquare(base float64) float64 {
	return base * base
}

func (this *Pixel) getZValues() {
	curIter := 0

	zReal0 := float64(this.X)*XScale - (FractalArea.XMax - FractalArea.XMin)
	zComplex0 := float64(this.Y)*YScale - (FractalArea.YMax - FractalArea.YMin)

	for ; (qsquare(this.ZReal)+qsquare(this.ZComplex) <= 4.0) && (curIter < MaxIters); curIter++ {
		nextZReal := qsquare(this.ZReal) - qsquare(this.ZComplex) + zReal0
		nextZComplex := 2*this.ZReal*this.ZComplex + zComplex0

		this.ZReal = nextZReal
		this.ZComplex = nextZComplex
	}

	this.getColor(curIter)
}

func (this Pixel) String() string {
	return fmt.Sprintf("\033[%sm \033[0m", *this.Color)
}

// TODO: limit logical threads to number of CPU threads.
// TODO: Make the cmdline arg parsing better.
// TODO: Print output lines in xSize sized batches, rather than one at a time.
func main() {
	fmt.Println("")
	cmdArgs := os.Args[1:]

	xSize, ex := strconv.Atoi(cmdArgs[0])
	ySize, ey := strconv.Atoi(cmdArgs[1])
	var ei error = nil // Fixes scoping problem for MaxIters
	MaxIters, ei = strconv.Atoi(cmdArgs[2])
	viewxmin, exmi := strconv.ParseFloat(cmdArgs[3], 64)
	viewxmax, exma := strconv.ParseFloat(cmdArgs[4], 64)
	viewymin, eymi := strconv.ParseFloat(cmdArgs[5], 64)
	viewymax, eyma := strconv.ParseFloat(cmdArgs[6], 64)

	// Much less useful than the previous error handling method, but allows for more
	for i, err := range []error{ex, ey, ei, exmi, exma, eymi, eyma} {
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Incorrect cmdline format at argument %d.\nExiting...\n", i+1)
			return
		}
	}

	FractalArea = ViewPort{viewxmin, viewxmax, viewymin, viewymax}

	pixels := make([]Pixel, xSize*ySize)
	// x/xtot = mx/mtot, where mtot = 4 (range of mandelbrot set)
	// -> x * mtot/xtot = mx
	// This gives us the scale, but not the exact value of zreal0.
	// Because the (min, max) of the mandelbrot == (-2.0, 2.0) for each variable,
	// we need to subtract 2 to get to the "true" scaled value.
	// Thus, we get:
	// zreal_0 = (x * mtot/xtot) - 2
	// This follows for y and zcomplex_0
	XScale = (FractalArea.XMax - FractalArea.XMin) / float64(xSize)
	YScale = (FractalArea.YMax - FractalArea.YMin) / float64(ySize)

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			index := x + y*xSize
			pixels[index] = Pixel{x, y, &ColorMap[0], 0.0, 0.0}
			waitgroup.Add(1)
			go func() {
				pixels[index].getZValues()
				waitgroup.Done()
			}()
		}
	}

	waitgroup.Wait()

	for _, pix := range pixels {
		fmt.Printf("%s", pix)
		if pix.X == xSize-1 {
			fmt.Print("\n")
		}
	}
}
