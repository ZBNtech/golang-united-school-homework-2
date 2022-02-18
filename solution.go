package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type mySide int

const SidesCircle mySide = 0
const SidesTriangle mySide = 3
const SidesSquare mySide = 4

func CalcSquare(sideLen float64, sidesNum mySide) float64 {

	var result float64

	if sidesNum == SidesCircle {
		result = CalcCircle(sideLen)
	} else if sidesNum == SidesTriangle {
		result = CalcTriangle(sideLen)
	} else if sidesNum == SidesSquare {
		result = CalcRectangle(sideLen)
	} else {
		result = 0.0
	}

	return result
}

func CalcRectangle(recLen float64) float64 {

	var RectangleSquare float64

	RectangleSquare = math.Pow(recLen, 2)

	return RectangleSquare
}

func CalcTriangle(triLen float64) float64 {

	var TriangleSquare float64

	TriangleSquare = (math.Sqrt(3) / 4) * math.Pow(triLen, 2)

	return TriangleSquare
}

func CalcCircle(cirLen float64) float64 {

	var CircleSquare float64

	CircleSquare = math.Pow(cirLen, 2) * math.Pi

	return CircleSquare
}
