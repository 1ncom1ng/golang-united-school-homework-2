package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sideNum int) float64 {
	var square float64
	const SidesTriangle = 3
	const SidesSquare = 4
	const SidesCircle = 0

	if sideNum == 0 {
		square = Pi * sideLen * sideLen
	} else if sideNum == 3 {
		square = math.Sqrt(3) * sideLen * sideLen / 4
	} else if sideNum == 4 {
		square = sideLen * sideLen
	} else {
		square = 0
	}

	return square
}
