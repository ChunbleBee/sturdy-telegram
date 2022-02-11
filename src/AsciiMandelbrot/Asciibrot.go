package main

import (
	"math"
)

type Pixel struct {
	X int
	Y int
	Color string
	ZReal float
	ZComplex float
}

var ColorMap = map[]

func getColor(iters int, maxIters int) {
	iterVal = math.Round((float(iters)/float(maxIters))*10)
}

func getZValues(pos Pixel, maxIters int) {
	curIter := 0
	for ; ((pos.ZReal * pos.ZReal) + (pos.ZComplex * pos.ZComplex) < 4.0) && (curIter < maxIter); curIter++ {
		nextZReal := (pos.ZReal * pos.ZReal) - (pos.ZComplex * pos.ZComplex) + float(pos.X)
		nextZComplex = 2 * pos.ZReal * pos.ZComplex + float(pos.Y)

		pos.ZReal = nextZReal
		pos.ZComplex = nextZComplex
	}
}

