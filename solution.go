package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intTypeForSquareCalculating int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

//CalcSquare
//Function for calculate square
func CalcSquare(sideLen float64, sidesNum intTypeForSquareCalculating) float64 {
	var square float64 = 0

	switch sidesNum {
	case SidesCircle:
		square = math.Pi * math.Pow(sideLen, 2)

	case SidesTriangle:
		square = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4

	case SidesSquare:
		square = math.Pow(sideLen, 2)
	}

	return square
}
