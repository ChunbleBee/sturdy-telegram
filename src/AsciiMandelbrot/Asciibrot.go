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

//var ColorMap = map[]
var maxIters int = 0;

func getColor(iters int, maxIters int) string {
	iterVal := int(math.Round((float64(iters)/float64(maxIters))*10))
	return string(iterVal)
}

func (this Pixel) getZValues(maxIters int) {
	curIter := 0

	for ; ((this.ZReal * this.ZReal) + (this.ZComplex * this.ZComplex) <= 4.0) && (curIter < maxIters); curIter++ {
		nextZReal := (this.ZReal * this.ZReal) - (this.ZComplex * this.ZComplex) + float64(this.X)
		nextZComplex := 2 * this.ZReal * this.ZComplex + float64(this.Y)

		this.ZReal = nextZReal
		this.ZComplex = nextZComplex
	}

	this.Color = getColor(curIter, maxIters)
}

func (this Pixel) String() string {
	return fmt.Sprintf("(%d, %d), (%f, %f), Color:%s ", this.X, this.Y, this.ZReal, this.ZComplex, this.Color)
}

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

	pixels := []Pixel{}

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			append(pixels, Pixel{x, y, "", 0.0, 0.0})
			go pixels[x + x*y].getZValues(iters)
		}
	}

	for pix := range pixels {
		fmt.Println(pix)
	}
}