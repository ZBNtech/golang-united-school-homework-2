package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myside int

func CalcSquare(sideLen float64, sidesNum myside) float64 {

	var result float64

	if sidesNum == 0 {
		result = CalcCircle(sideLen)
	} else if sidesNum == 3 {
		result = CalcTriangle(sideLen)
	} else if sidesNum == 4 {
		result = CalcRectangle(sideLen)
	} else {
		result = 0.0
	}
	
	return result
}

func CalcRectangle(sideLen float64) float64 {

	var RectangleSquare float64

	RectangleSquare = sideLen * 2

	return RectangleSquare
}

func CalcTriangle(sideLen float64) float64 {

	var TriangleSquare float64

	TriangleSquare = (sideLen * 2 * math.Sqrt(3.0)) / 4

	return TriangleSquare
}

func CalcCircle(sideLen float64) float64 {

	var CircleSquare float64

	CircleSquare = sideLen * 2 * math.Pi

	return CircleSquare
}
