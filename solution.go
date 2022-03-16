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

type numberSides int;

func CalcSquare(sideLen float64, sidesNum numberSides) float64 {
	var square float64
	const SidesTriangle numberSides = 3
	const SidesSquare numberSides = 4
	const SidesCircle numberSides = 0

	if sidesNum == SidesCircle {
		square = Pi * sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		square = math.Sqrt(3) * sideLen * sideLen / 4
	} else if sidesNum == SidesSquare {
		square = sideLen * sideLen
	} else {
		square = 0
	}

	return square
}
